PROTOS_LOC:=v2/protos
PROTOBUF_BUILDER_IMG:=pbnj-protobuf-builder

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: pbs
pbs: ## locally generate go stubs from protocol buffers
	scripts/protoc.sh

.PHONY: pbs-docker-image
pbs-docker-image: ## generate container image for building protocol buffers 
	docker build -t ${PROTOBUF_BUILDER_IMG} -f scripts/Dockerfile.pbbuilder .

.PHONY: pbs-docker
pbs-docker: pbs-docker-image ## generate go stubs from protocol buffers in a container
	docker run -it --rm -v ${PWD}:/code -w /code ${PROTOBUF_BUILDER_IMG} scripts/protoc.sh

.PHONY: evans
evans: ## run evans grpc client
	evans -p 60061 -r repl
