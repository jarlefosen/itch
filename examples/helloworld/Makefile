.PHONY: all build-scratch build-itch build clean run-scratch run-itch run

IMG_ITCH := example-helloworld:itch
IMG_SCRATCH := example-helloworld:scratch

all: build run

build-scratch:
	docker build -t $(IMG_SCRATCH) -f Dockerfile.scratch .

build-itch:
	docker build -t $(IMG_ITCH) -f Dockerfile .

build: build-itch build-scratch

clean:
	docker rmi $(IMG_SCRATCH)
	docker rmi $(IMG_ITCH)

run-scratch:
	@echo Running $(IMG_SCRATCH)
	@docker run --rm $(IMG_SCRATCH)

run-itch:
	@echo Running $(IMG_ITCH)
	@docker run --rm $(IMG_ITCH)

run: run-scratch run-itch
