VERSION := $(shell git describe --tags --dirty=-modified --always)
SERVICE_CONTROLLER_IMAGE := quay.io/skupper/service-controller
CONTROLLER_PODMAN_IMAGE := quay.io/skupper/controller-podman
SITE_CONTROLLER_IMAGE := quay.io/skupper/site-controller
CONFIG_SYNC_IMAGE := quay.io/skupper/config-sync
FLOW_COLLECTOR_IMAGE := quay.io/skupper/flow-collector
TEST_IMAGE := quay.io/skupper/skupper-tests
TEST_BINARIES_FOLDER := ${PWD}/test/integration/bin
DOCKER := docker
LDFLAGS := -X github.com/skupperproject/skupper/pkg/version.Version=${VERSION}

all: generate-client build-cmd build-get build-config-sync build-controllers build-tests build-manifest

build-tests:
	mkdir -p ${TEST_BINARIES_FOLDER}
	go test -c -tags=job -v ./test/integration/examples/tcp_echo/job -o ${TEST_BINARIES_FOLDER}/tcp_echo_test
	go test -c -tags=job -v ./test/integration/examples/http/job -o ${TEST_BINARIES_FOLDER}/http_test
	go test -c -tags=job -v ./test/integration/examples/bookinfo/job -o ${TEST_BINARIES_FOLDER}/bookinfo_test
	go test -c -tags=job -v ./test/integration/examples/mongodb/job -o ${TEST_BINARIES_FOLDER}/mongo_test
	go test -c -tags=job -v ./test/integration/examples/custom/hipstershop/job -o ${TEST_BINARIES_FOLDER}/grpcclient_test
	go test -c -tags=job -v ./test/integration/examples/tls_t/job -o ${TEST_BINARIES_FOLDER}/tls_test

build-cmd: generate-client
	go build -ldflags="${LDFLAGS}"  -o skupper ./cmd/skupper

build-get:
	go build -ldflags="${LDFLAGS}"  -o get ./cmd/get

build-service-controller:
	go build -ldflags="${LDFLAGS}"  -o service-controller cmd/service-controller/main.go cmd/service-controller/controller.go cmd/service-controller/ports.go cmd/service-controller/definition_monitor.go cmd/service-controller/console_server.go cmd/service-controller/site_query.go cmd/service-controller/ip_lookup.go cmd/service-controller/token_handler.go cmd/service-controller/secret_controller.go cmd/service-controller/claim_handler.go cmd/service-controller/tokens.go cmd/service-controller/links.go cmd/service-controller/services.go cmd/service-controller/policies.go cmd/service-controller/policy_controller.go cmd/service-controller/revoke_access.go  cmd/service-controller/nodes.go

build-controller-podman:
	go build -ldflags="${LDFLAGS}"  -o controller-podman cmd/controller-podman/main.go

build-site-controller:
	go build -ldflags="${LDFLAGS}"  -o site-controller cmd/site-controller/main.go cmd/site-controller/controller.go

build-flow-collector:
	go build -ldflags="${LDFLAGS}"  -o flow-collector cmd/flow-collector/main.go cmd/flow-collector/controller.go cmd/flow-collector/handlers.go

build-config-sync:
	go build -ldflags="${LDFLAGS}"  -o config-sync cmd/config-sync/main.go cmd/config-sync/config_sync.go cmd/config-sync/collector.go

build-controllers: build-site-controller build-service-controller build-controller-podman build-flow-collector

build-manifest:
	go build -ldflags="${LDFLAGS}"  -o manifest ./cmd/manifest

docker-build-test-image:
	${DOCKER} build -t ${TEST_IMAGE} -f Dockerfile.ci-test .

docker-build: generate-client docker-build-test-image
	${DOCKER} build -t ${SERVICE_CONTROLLER_IMAGE} -f Dockerfile.service-controller .
	${DOCKER} build -t ${CONTROLLER_PODMAN_IMAGE} -f Dockerfile.controller-podman .
	${DOCKER} build -t ${SITE_CONTROLLER_IMAGE} -f Dockerfile.site-controller .
	${DOCKER} build -t ${CONFIG_SYNC_IMAGE} -f Dockerfile.config-sync .
	${DOCKER} build -t ${FLOW_COLLECTOR_IMAGE} -f Dockerfile.flow-collector .

docker-push-test-image:
	${DOCKER} push ${TEST_IMAGE}

docker-push: docker-push-test-image
	${DOCKER} push ${SERVICE_CONTROLLER_IMAGE}
	${DOCKER} push ${CONTROLLER_PODMAN_IMAGE}
	${DOCKER} push ${SITE_CONTROLLER_IMAGE}
	${DOCKER} push ${CONFIG_SYNC_IMAGE}
	${DOCKER} push ${FLOW_COLLECTOR_IMAGE}

format:
	go fmt ./...

generate-client:
	./scripts/update-codegen.sh
	./scripts/libpod-generate.sh

force-generate-client:
	FORCE=true ./scripts/update-codegen.sh
	FORCE=true ./scripts/libpod-generate.sh

client-mock-test:
	go test -v -count=1 ./client

client-cluster-test:
	go test -v -count=1 ./client -use-cluster

vet:
	go vet ./...

cmd-test:
	go test -v -count=1 ./cmd/...

pkg-test:
	go test -v -count=1 ./pkg/...

.PHONY: test
test:
	go test -v -count=1 ./pkg/... ./cmd/... ./client/...

clean:
	rm -rf skupper service-controller controller-podman site-controller release get config-sync manifest ${TEST_BINARIES_FOLDER}

package: release/windows.zip release/darwin.zip release/linux.tgz

release/linux.tgz: release/linux/skupper
	tar -czf release/linux.tgz -C release/linux/ skupper

release/linux/skupper: cmd/skupper/skupper.go
	GOOS=linux GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o release/linux/skupper ./cmd/skupper

release/windows/skupper: cmd/skupper/skupper.go
	GOOS=windows GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o release/windows/skupper ./cmd/skupper

release/windows.zip: release/windows/skupper
	zip -j release/windows.zip release/windows/skupper

release/darwin/skupper: cmd/skupper/skupper.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o release/darwin/skupper ./cmd/skupper

release/darwin.zip: release/darwin/skupper
	zip -j release/darwin.zip release/darwin/skupper

generate-manifest: build-manifest
	./manifest

# If the control pane api changes, you can update the api client by running:
#    $ cp ../nexodus/internal/docs/swagger.yaml internal/control-api/swagger.yaml && make generate-openapi-client
.PHONY: generate-openapi-client
generate-openapi-client: ./internal/control-api/client.go ## Generate the OpenAPI Client for the control-api
internal/control-api/%.go: internal/control-api/client.go
internal/control-api/client.go: ./internal/control-api/swagger.yaml
	rm -f $(shell find ./internal/control-api | grep .go | grep -v custom_)
	docker run --rm -v $(CURDIR):/src --user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:v6.5.0 \
		generate -i /src/internal/control-api/swagger.yaml -g go \
		--package-name control_api \
		-o /src/internal/control-api \
		--ignore-file-override /src/.openapi-generator-ignore
	gofmt -w ./internal/control-api

