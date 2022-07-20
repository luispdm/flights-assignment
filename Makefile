export DATA ?= testdata/long.json
COMPOSE	   := docker/docker-compose.yml

.PHONY: cli
cli:
	@ cd cmd/cli && \
		go run main.go

bench:
	@ cd cmd/cli && \
		go test \
			-bench=. \
			-benchmem
http:
	@ go run ./cmd/http/.

track:
	@ curl \
		-i \
		-X POST \
		-d @${DATA} \
		http://localhost:8080/track

health:
	@ curl \
		-i \
		http://localhost:8080/health

binary:
	@ CGO_ENABLED=0 \
		GOOS=linux \
		go build \
			-v \
			-o \
			server \
			./cmd/http/.

dev:
	@ docker-compose \
		-f \
		${COMPOSE} \
		up \
		-d

logs:
	@ docker-compose \
		-f \
		${COMPOSE} \
		logs \
		-f

nodev:
	@ docker-compose \
		-f \
		${COMPOSE} \
		down

devrebuild:
	@ docker-compose \
		-f \
		${COMPOSE} \
		up \
		--build \
		--remove-orphans \
		-d
