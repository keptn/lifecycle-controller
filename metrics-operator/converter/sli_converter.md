# SLI Converter

SLI converter is a tool to convert the old `sli.yaml`
[file format](https://github.com/keptn/spec/blob/master/service_level_indicator.md) into the new
`AnalysisValueTemplate` [custom resource definition](https://keptn.sh/latest/docs/reference/api-reference/metrics/v1/).
The converter is part of `metrics-operator` image.

## Usage

The converter will convert a single `sli.yaml` file into multiple `AnalysisValueTemplate` resources.

To run the converter, execute the following command:

<!---x-release-please-start-version-->
```shell
METRICS_OPERATOR_IMAGE=ghcr.io/keptn/metrics-operator:v2.0.1
PATH_TO_SLI=<PATH_TO_SLI>
KEPTN_PROVIDER_NAME=<KEPTN_PROVIDER_NAME>
KEPTN_PROVIDER_NAMESPACE=<KEPTN_PROVIDER_NAMESPACE>

docker run $METRICS_OPERATOR_IMAGE manager --convert-sli=$PATH_TO_SLI --keptn-provider-name=$KEPTN_PROVIDER_NAME --keptn-provider-namespace=$KEPTN_PROVIDER_NAMESPACE
```
<!---x-release-please-end-->

Please be aware, you need to substitute the placeholders with the following information:

* **PATH_TO_SLI** - path to your `sli.yaml` file
* **KEPTN_PROVIDER_NAME** - name of `KeptnMetricsProvider` which will be used to fetch SLIs
* **KEPTN_PROVIDER_NAMESPACE** - namespace of `KeptnMetricsProvider` which will be used to fetch SLIs

> **Note**

All the SLIs present in `sli.yaml` file will use the same provider defined by the referenced
`KeptnMetricsProvider`.

## Example

The following content of `sli.yaml` file

```yaml
spec_version: "1.0"
indicators:
  throughput: "builtin:service.requestCount.total:merge(0):count?scope=tag(keptn_project:$PROJECT),tag(keptn_stage:$STAGE),tag(keptn_service:$SERVICE),tag(keptn_deployment:$DEPLOYMENT)"
  response_time_p95: "builtin:service.response.time:merge(0):percentile(95)?scope=tag(keptn_project:$PROJECT),tag(keptn_stage:$STAGE),tag(keptn_service:$SERVICE),tag(keptn_deployment:$DEPLOYMENT)"
```

with the following command

```shell
docker run $METRICS_OPERATOR_IMAGE manager --convert-sli=./sli.yaml --sli-provider=dynatrace --sli-namespace=keptn
```

will be converted to:

```yaml
---
apiVersion: metrics.keptn.sh/v1
kind: AnalysisValueTemplate
metadata:
  creationTimestamp: null
  name: response_time_p95
spec:
  provider:
    name: dynatrace
    namespace: keptn
  query: builtin:service.response.time:merge(0):percentile(95)?scope=tag(keptn_project:{{.project}}),tag(keptn_stage:{{.stage}}),tag(keptn_service:{{.service}}),tag(keptn_deployment:{{.deployment}})
---
apiVersion: metrics.keptn.sh/v1
kind: AnalysisValueTemplate
metadata:
  creationTimestamp: null
  name: throughput
spec:
  provider:
    name: dynatrace
    namespace: keptn
  query: builtin:service.requestCount.total:merge(0):count?scope=tag(keptn_project:{{.project}}),tag(keptn_stage:{{.stage}}),tag(keptn_service:{{.service}}),tag(keptn_deployment:{{.deployment}})
```
