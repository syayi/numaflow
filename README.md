# Numaflow

[![Go Report Card](https://goreportcard.com/badge/github.com/numaproj/numaflow)](https://goreportcard.com/report/github.com/numaproj/numaflow)
[![slack](https://img.shields.io/badge/slack-numaproj-brightgreen.svg?logo=slack)](https://join.slack.com/t/numaproj/shared_invite/zt-19svuv47m-YKHhsQ~~KK9mBv1E7pNzfg)
[![GoDoc](https://godoc.org/github.com/numaproj/numaflow?status.svg)](https://godoc.org/github.com/numaproj/numaflow/pkg/apis)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Release Version](https://img.shields.io/github/v/release/numaproj/numaflow?label=numaflow)](https://github.com/numaproj/numaflow/releases/latest)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/6078/badge)](https://bestpractices.coreinfrastructure.org/projects/6078)

## Summary

Numaflow is a Kubernetes-native platform for developing event-driven applications and event processing. A Numaflow Pipeline is implemented
as a Kubernetes custom resource and consists of one or more sources, processing, and sink vertices.

Numaflow installs in a few minutes and is easier and cheaper to use for developing event-driven applications.

## Use Cases

- Event-driven applications such as anomaly detection, monitoring, and alerting.
- Real-time analytics.
- Streaming applications such as data instrumentation and data movement.
- Workflows running in a streaming manner.

## Key Features

- Kubernetes-native: If you know Kubernetes, you already know how to use Numaflow.
- Language agnostic: Use your favorite programming language.
- Auto-scaling with back-pressure: Each vertex automatically scales from zero to whatever is needed.
- Ready-to-use sources/sinks: Kafka, Pulsar, HTTP, and more. Easy to write custom sources and sinks.
- Exactly-Once semantics: No input element is duplicated or lost even as pods are rescheduled or restarted.

## Data Integrity Guarantees

- Minimally provide at-least-once semantics
- Provide exactly-once semantics for unbounded and near real-time data sources
- Preserving order is not required

## Demo

[![Numaflow Demo](https://img.youtube.com/vi/TOqKOYX0nrE/0.jpg)](https://youtu.be/TOqKOYX0nrE)

## Resources

- [QUICK_START](docs/quick-start.md)
- [EXAMPLES](examples)
- [DEVELOPMENT](docs/development/development.md)
- [CONTRIBUTING](https://github.com/numaproj/numaproj/blob/main/CONTRIBUTING.md)

## Get Support
There are several ways to get in touch with us:
* [Slack Community](https://numaproj.slack.com/join/shared_invite/zt-19svuv47m-YKHhsQ~~KK9mBv1E7pNzfg#/shared-invite/email)

## Roadmap

- Map Streaming (1.3)
