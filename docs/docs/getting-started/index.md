---
comments: true
---

# Get started

This section provides tutorials to familiarize you
with some basic Keptn features
and point you to other documentation that has more comprehensive information.

This section provides tutorials to familiarize you
with some basic Keptn features
and point you to other documentation that has more comprehensive information.

Before attempting these tutorials, we recommend that you read the
[Core Concepts]((../core-concepts/index.md)) section.

## Keptn Metrics

The Keptn Metrics component simplifies observability in Kubernetes
by unifying and standardizing metrics from multiple data sources,
such as Prometheus, Dynatrace, Datadog, and cloud providers like AWS, Google,
and Azure.
It integrates seamlessly with deployment and scaling tools like Argo,
Flux, KEDA, and HPA, enabling automated decisions with minimal configuration.
Unlike the Kubernetes Metrics Server, Keptn Metrics eliminates the complexity
of maintaining multiple point-to-point integrations, offering a streamlined,
centralized approach for managing metrics across diverse tools and observability platforms.

## Keptn Observability

Keptn enhances your cloud-native deployment environment with powerful observability
features, whether or not you follow a GitOps strategy.
This guide walks you through installing Keptn on a Kubernetes cluster, enabling
Keptn for a namespace and deployment, and setting up Grafana and observability tools to
visualize DORA metrics and OpenTelemetry traces.
By the end, you’ll have a fully integrated observability system built step-by-step.

## Release Lifecycle Management

Keptn's Release Lifecycle Management tools enhance your cloud-native deployments by
running pre- and post-deployment tasks, conducting SLO evaluations, and managing
workload promotions.
This tutorial introduces these tools and demonstrates how to configure Keptn to trigger
webhooks around deployments.
It builds on the **Getting Started with Keptn Observability** exercise, so completing
that first is recommended.
You'll learn to annotate workloads, group them into KeptnApp resources, and set
up webhook triggers for robust deployment workflows.
