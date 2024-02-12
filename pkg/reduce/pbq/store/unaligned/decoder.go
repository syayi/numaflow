package unaligned

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/numaproj/numaflow/pkg/isb"
	"github.com/numaproj/numaflow/pkg/reduce/pbq/partition"
)

const (
	dMessageHeaderSize = 22
	EntryHeaderSize    = 28
)

var (
	location            *time.Location
	errChecksumMismatch = fmt.Errorf("data checksum not match")
)

func init() {
	var err error
	location, err = time.LoadLocation("UTC")
	if err != nil {
		panic(fmt.Sprint("cannot load UTC", err))
	}
}

// decoder is a decoder for the WAL entries and header.
type decoder struct{}

// newDecoder returns a new decoder
func newDecoder() *decoder {
	return &decoder{}
}

// decodeHeader decodes the header from the given io.Reader.
// the header of the WAL file is a partition.ID, so it returns a partition.ID.
func (d *decoder) decodeHeader(buf io.Reader) (*partition.ID, error) {
	var err error

	// read the fixed values
	var hp = new(walHeaderPreamble)
	err = binary.Read(buf, binary.LittleEndian, hp)
	if err != nil {
		return nil, err
	}

	// read the variadic slot
	var slot = make([]rune, hp.SLen)
	err = binary.Read(buf, binary.LittleEndian, slot)
	if err != nil {
		return nil, err
	}

	return &partition.ID{
		Start: time.UnixMilli(hp.S).In(location),
		End:   time.UnixMilli(hp.E).In(location),
		Slot:  string(slot),
	}, nil
}

// decodeMessage decodes the isb read message from the given io.Reader.
func (d *decoder) decodeMessage(buf io.Reader) (*isb.ReadMessage, string, error) {
	entryHeader, err := d.decodeWALMessageHeader(buf)
	if err != nil {
		return nil, "", err
	}

	// read the key
	key := make([]rune, entryHeader.KeyLen)
	err = binary.Read(buf, binary.LittleEndian, &key)
	if err != nil {
		return nil, "", err
	}

	entryBody, err := d.decodeWALBody(buf, entryHeader)
	if err != nil {
		return nil, "", err
	}

	return &isb.ReadMessage{
		Message:    *entryBody,
		Watermark:  time.UnixMilli(entryHeader.WaterMark).In(location),
		ReadOffset: isb.SimpleIntOffset(func() int64 { return entryHeader.Offset }),
	}, string(key), nil
}

// decodeDeletionMessage decodes deletion message from the given io.Reader
func (d *decoder) decodeDeletionMessage(buf io.Reader) (*deletionMessage, int64, error) {
	dms := deletionMessage{}

	dMessageHeader := deletionMessageHeaderPreamble{}
	if err := binary.Read(buf, binary.LittleEndian, &dMessageHeader); err != nil {
		return nil, 0, err
	}

	dms.St = dMessageHeader.St
	dms.Et = dMessageHeader.Et

	// read the slot
	var slot = make([]rune, dMessageHeader.SLen)
	if err := binary.Read(buf, binary.LittleEndian, slot); err != nil {
		return nil, 0, err
	}

	dms.Slot = string(slot)

	// read the key
	var key = make([]rune, dMessageHeader.KLen)
	if err := binary.Read(buf, binary.LittleEndian, key); err != nil {
		return nil, 0, err
	}

	dms.Key = string(key)

	size := dMessageHeaderSize + int64(dMessageHeader.SLen) + int64(dMessageHeader.KLen)
	return &dms, size, nil
}

// decodeWALMessageHeader decodes the WAL message header from the given io.Reader.
func (d *decoder) decodeWALMessageHeader(buf io.Reader) (*readMessageHeaderPreamble, error) {
	var entryHeader = new(readMessageHeaderPreamble)
	err := binary.Read(buf, binary.LittleEndian, entryHeader)
	if err != nil {
		return nil, err
	}
	return entryHeader, nil
}

// decodeWALBody decodes the WAL message body from the given io.Reader.
func (d *decoder) decodeWALBody(buf io.Reader, entryHeader *readMessageHeaderPreamble) (*isb.Message, error) {
	var err error

	body := make([]byte, entryHeader.MessageLen)
	size, err := buf.Read(body)
	if err != nil {
		return nil, err
	}
	if int64(size) != entryHeader.MessageLen {
		return nil, fmt.Errorf("expected to read length of %d, but wrote only %d", entryHeader.MessageLen, size)
	}
	// verify the checksum
	checksum := calculateChecksum(body)
	if checksum != entryHeader.Checksum {
		println("header - ", entryHeader.MessageLen, " ", entryHeader.Checksum, " ", entryHeader.KeyLen, " ", entryHeader.Offset, " ", entryHeader.WaterMark, " ", entryHeader.EventTime)
		println("body - ", string(body))
		var message = new(isb.Message)
		err = message.UnmarshalBinary(body)
		if err != nil {
			println("error - ", err.Error())
		}
		return nil, errChecksumMismatch
	}

	var message = new(isb.Message)
	err = message.UnmarshalBinary(body)
	if err != nil {
		return nil, err
	}
	return message, nil
}
