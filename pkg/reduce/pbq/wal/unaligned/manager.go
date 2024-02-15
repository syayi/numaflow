/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package unaligned

import (
	"context"
	"os"
	"strings"
	"time"

	dfv1 "github.com/numaproj/numaflow/pkg/apis/numaflow/v1alpha1"
	"github.com/numaproj/numaflow/pkg/reduce/pbq/partition"
	"github.com/numaproj/numaflow/pkg/reduce/pbq/wal"
)

type fsWAL struct {
	storePath               string
	pipelineName            string
	vertexName              string
	replicaIndex            int32
	fsOpts                  []WALOption
	segmentRotationDuration time.Duration
	activeWals              map[string]wal.WAL
}

// NewFSManager is a FileSystem Stores Manager.
func NewFSManager(storePath string, vertexInstance *dfv1.VertexInstance, opts ...WALOption) wal.Manager {
	s := &fsWAL{
		storePath:    storePath,
		pipelineName: vertexInstance.Vertex.Spec.PipelineName,
		vertexName:   vertexInstance.Vertex.Spec.AbstractVertex.Name,
		replicaIndex: vertexInstance.Replica,
		activeWals:   make(map[string]wal.WAL),
		fsOpts:       opts,
	}
	return s
}

// CreateWAL creates the FS unalignedWAL.
func (ws *fsWAL) CreateWAL(_ context.Context, partitionID partition.ID) (wal.WAL, error) {
	// check if the wal is already present
	// during crash recovery, we might have already created the wal while replaying
	if store, ok := ws.activeWals[partitionID.String()]; ok {
		return store, nil
	}
	// Create fs dir if not exist
	var err error
	if _, err = os.Stat(ws.storePath); os.IsNotExist(err) {
		err = os.Mkdir(ws.storePath, 0755)
		if err != nil {
			return nil, err
		}
	}

	w, err := NewUnalignedWriteOnlyWAL(&partitionID, ws.fsOpts...)
	if err != nil {
		return nil, err
	}
	ws.activeWals[w.PartitionID().String()] = w
	return w, nil
}

// DiscoverWALs returns all the wals present in the storePath
func (ws *fsWAL) DiscoverWALs(_ context.Context) ([]wal.WAL, error) {
	files, err := os.ReadDir(ws.storePath)
	if os.IsNotExist(err) {
		return []wal.WAL{}, nil
	} else if err != nil {
		return nil, err
	}
	partitions := make([]wal.WAL, 0)

	for _, f := range files {
		if strings.HasPrefix(f.Name(), SegmentPrefix) && !f.IsDir() {
			wl, err := NewUnalignedReadWriteWAL(ws.fsOpts...)
			if err != nil {
				return nil, err
			}
			partitions = append(partitions, wl)
			// add to activeWals, to avoid creating the same wal again
			// after replay.
			ws.activeWals[wl.PartitionID().String()] = wl
		}
	}

	return partitions, nil
}

// DeleteWAL deletes the store for the given partitionID
func (ws *fsWAL) DeleteWAL(partitionID partition.ID) error {
	w, ok := ws.activeWals[partitionID.String()]
	if !ok {
		return nil
	}
	delete(ws.activeWals, partitionID.String())
	return w.Close()
}
