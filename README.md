# Numaflow

[![Go Report Card](https://goreportcard.com/badge/github.com/numaproj/numaflow)](https://goreportcard.com/report/github.com/numaproj/numaflow)
[![slack](https://img.shields.io/badge/slack-numaproj-brightgreen.svg?logo=slack)](https://join.slack.com/t/numaproj/shared_invite/zt-19svuv47m-YKHhsQ~~KK9mBv1E7pNzfg)
[![GoDoc](https://godoc.org/github.com/numaproj/numaflow?status.svg)](https://godoc.org/github.com/numaproj/numaflow/pkg/apis)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Release Version](https://img.shields.io/github/v/release/numaproj/numaflow?label=numaflow)](https://github.com/numaproj/numaflow/releases/latest)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/6078/badge)](https://bestpractices.coreinfrastructure.org/projects/6078)

Welcome to Numaflow, a Kubernetes-native platform designed for developing scalable and reliable event-driven applications. Numaflow decouples event sources and sinks from the processing logic, allowing each component to independently auto-scale based on demand. With out-of-the-box sources and sinks, developers can focus solely on the processing logic without worrying about event consumption, boilerplate code, or operational complexities. Each step of the pipeline can be written in any programming language, offering unparalleled flexibility and ease of use

Numaflow, created by the team behind Argo, leverages their expertise to deliver a robust and scalable end-to-end platform for event-driven applications

![Numaflow Pipeline](./ui/src/images/simple-pipeline.png)

## Use Cases

- Event driven applications: Process events as they happen, e.g., updating inventory and sending customer notifications in e-commerce
- Real time analytics: Analyze data instantly, e.g., social media analytics, observability data processing
- Inference on streaming data: Perform real-time predictions, e.g., anomaly detection
- Workflows running in a streaming manner

## Key Features

- Kubernetes-native: If you know Kubernetes, you already know how to use Numaflow
- Language agnostic: Use your favorite programming language
- Exactly-Once semantics: No input element is duplicated or lost even as pods are rescheduled or restarted
- Auto-scaling with back-pressure: Each vertex automatically scales from zero to whatever is needed

## Data Integrity Guarantees

- Minimally provide at-least-once semantics
- Provide exactly-once semantics for unbounded and near real-time data sources
- Preserving order is not required

## Roadmap

- Map Streaming (1.3)

## Demo

[![Numaflow Demo](https://img.youtube.com/vi/TOqKOYX0nrE/0.jpg)](https://youtu.be/TOqKOYX0nrE)

## Resources

- [QUICK_START](docs/quick-start.md)
- [EXAMPLES](examples)
- [DEVELOPMENT](docs/development/development.md)
- [CONTRIBUTING](https://github.com/numaproj/numaproj/blob/main/CONTRIBUTING.md)
