
# Image URL to use all building/pushing image targets
IMG ?= certificate-operator:latest
# RELEASE_REGISTRY is the container registry to push into.
RELEASE_REGISTRY?=ghcr.io/keptn
RELEASE_TIME=$(shell date +%Y%m%d%s)
BUILD_TIME=$(shell date -u "+%F_%T")
RELEASE_VERSION?=$(RELEASE_TIME)-v0.24.3#$(shell git describe --tags --match "v*")
TAG?="v3.0.1" # x-release-please-version
RELEASE_IMAGE:=certificate-operator:$(TAG)
CHART_APPVERSION?=""

ARCHS = amd64 arm64
COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z)
BUILDENVVAR=CGO_ENABLED=0

VERSION=$(shell echo $(RELEASE_VERSION) | awk -F - '{print $$2}')

# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.27.0

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

#### developer build   ####
.PHONY: clean
clean:
	rm -rf ./bin
##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: manifests
manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) rbac:roleName=certificate-operator-role crd paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: unit-test
unit-test: manifests generate fmt vet envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(LOCALBIN) -p path)" go test ./... -coverprofile cover.out

#command(make test) to run all tests 
.PHONY: test
test: unit-test
##@ Build

.PHONY: build
build: generate fmt vet ## Build manager binary.
	go build -o bin/manager main.go

.PHONY: run
run: manifests generate fmt vet ## Run a controller from your host.
	go run ./main.go

##@ Deployment

ifndef ignore-not-found
  ignore-not-found = false
endif

.PHONY: deploy
deploy: ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && kustomize edit set image controller=${IMG}
	kustomize build config/default | kubectl apply -f -

.PHONY: undeploy
undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	kustomize build config/default | kubectl delete --ignore-not-found=$(ignore-not-found) -f -

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
KUSTOMIZE ?= $(LOCALBIN)/kustomize
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
ENVTEST ?= $(LOCALBIN)/setup-envtest

## Tool Versions

# renovate: datasource=github-tags depName=kubernetes-sigs/kustomize
KUSTOMIZE_VERSION?=v5.5.0
# renovate: datasource=github-releases depName=kubernetes-sigs/controller-tools
CONTROLLER_TOOLS_VERSION?=v0.16.5

KUSTOMIZE_INSTALL_SCRIPT ?= "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"
.PHONY: kustomize
kustomize: $(KUSTOMIZE) ## Download kustomize locally if necessary.
$(KUSTOMIZE): $(LOCALBIN)
	test -s $(LOCALBIN)/kustomize || { curl -Ss $(KUSTOMIZE_INSTALL_SCRIPT) | bash -s -- $(subst v,,$(KUSTOMIZE_VERSION)) $(LOCALBIN); }

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

.PHONY: envtest
envtest: $(ENVTEST) ## Download envtest-setup locally if necessary.
$(ENVTEST): $(LOCALBIN)
	test -s $(LOCALBIN)/setup-envtest || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest

.PHONY: release-local
release-local: release-local.amd64 release-local.arm64
	for arch in $(ARCHS); do \
			docker push $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-$${arch} ;\
		done
		DOCKER_CLI_EXPERIMENTAL=enabled docker manifest create $(RELEASE_REGISTRY)/$(RELEASE_IMAGE) $(addprefix --amend $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-, $(ARCHS))
		for arch in $(ARCHS); do \
			DOCKER_CLI_EXPERIMENTAL=enabled docker manifest annotate --arch $${arch} $(RELEASE_REGISTRY)/$(RELEASE_IMAGE) $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-$${arch} ;\
		done
		DOCKER_CLI_EXPERIMENTAL=enabled docker manifest push $(RELEASE_REGISTRY)/$(RELEASE_IMAGE) ;\

.PHONY: release-local.amd64
release-local.amd64: clean
	DOCKER_BUILDKIT=1 docker build \
		--build-arg GIT_HASH="$(HASH)" \
		--build-arg RELEASE_VERSION="$(TAG)" \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		--build-arg CONTROLLER_TOOLS_VERSION="$(CONTROLLER_TOOLS_VERSION)" \
		-t $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-amd64 \
		--platform linux/amd64 \
		--target debug \
		.
	docker tag $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-amd64 $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)

.PHONY: release-local.arm64
release-local.arm64: clean
	DOCKER_BUILDKIT=1 docker build \
		--build-arg GIT_HASH="$(HASH)" \
		--build-arg RELEASE_VERSION="$(TAG)" \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		--build-arg CONTROLLER_TOOLS_VERSION="$(CONTROLLER_TOOLS_VERSION)" \
		-t $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-arm64 \
		--platform linux/arm64 \
		--target debug \
		.
	docker tag $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-arm64 $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)

.PHONY: push-local
push-local:
	if [[ "${ARCH}" == "amd64" ]]; then \
		docker push $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-amd64 ;\
	elif [[ "${ARCH}" == "arm64" ]]; then \
		docker push $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)-arm64 ;\
	fi
	docker push $(RELEASE_REGISTRY)/$(RELEASE_IMAGE)

.PHONY: release-manifests
release-manifests:
	kustomize build config/default > config/rendered/release.yaml
	envsubst < config/rendered/release.yaml > tmp.yaml; mv tmp.yaml config/rendered/release.yaml

##golangci-lint
.PHONY: lint
lint:
	${GOPATH}/bin/golangci-lint run --config ../.golangci.yml -v
