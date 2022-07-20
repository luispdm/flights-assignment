# Flights assignment

Solution to the take-home assignment of an interview process I took part in 2022. Description of the assignment can be found [here](problem/README.md).

## System spec
All the work has been done on a 2021 16" Apple Macbook Pro with the following specs:
- M1 Pro processor (10-core CPU and 16-core GPU)
- 32 GB of RAM
- 1 TB SSD
### Software
- Go 1.18.3 (with [chi](https://github.com/go-chi/chi/) 5.0.7)
- macOS Monterey 12.4

## Considerations

### Algorithms and benchmarks
Before developing the server, some algorithms were developed and benchmarked. A compromise between speed, efficiency, and readability led to the choice of the algorithm used in the server.

Algorithms and benchmarks can be found under [cmd/cli](./cmd/cli/).

Benchmark results:
![Benchmark results](/img/bench_results.jpg)

The algorithm chosen is [`sliceDiff`](./cmd/cli/main.go).

### Sorting and non-duplicates
As far as I understood, sorting the list of flights was a suggestion, not a requirement. All the algorithms implemented don't sort the flights: they look for non-duplicated elements.

[`Track()`](./internal/tracker/tracker.go) contains the algorithm used by the server.

## API Format
The API has two endpoints:
- [`/health`](#health)
- [`/track`](#track)

### `/health`
Basic health check to detect if the application is responding. It answers to `GET` requests only.

#### Request example
```bash
curl -i http://localhost:8080
```

#### Response example
```bash
HTTP/1.1 200 OK
Date: Sat, 09 Jul 2022 17:51:45 GMT
Content-Length: 0
```

### `/track`
It returns the starting airport and the final airport from a list of unordered flights.

It answers to `POST` requests only. The list of flights should be fed via `.json`.

#### Json example
```json
[
    {
        "start": "SFO",
        "end": "ATL"
    },
    {
        "start": "ATL",
        "end": "GSO"
    }
]
```

More examples can be found under [testdata](./testdata/).

#### Request example
```bash
curl -i \
    -X \
    POST \
    -d \
    @testdata/long.json \
    http://localhost:8080/track
```

#### Response example
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 09 Jul 2022 18:01:21 GMT
Content-Length: 36

{
  "start": "BGY",
  "end": "AKL"
}
```

## How to run the server
**Prerequisite**: you need either Golang or Docker installed on your machine.

The server runs on port `8080`.

### Go
Before running the server you might want to execute `go mod tidy` to download potentially missing dependencies.

- `make http` runs the server on your machine

### Docker
The docker-compose can be found [here](./docker/docker-compose.yml). The docker-compose builds the Docker image corresponding to [this](docker/server/Dockerfile) Dockerfile and starts the server.

- `make dev` runs the docker-compose in the background
- `make logs` follows the docker-compose logs
- `make nodev` stops and removes the container
- `make devrebuild` rebuilds the image and runs the container in the background

## Execute requests
The [Makefile](./Makefile) also contains commands to execute HTTP requests to the server's endpoints.

- `make health` performs a `GET` request to the `/health` endpoint
- `make track` performs a `POST` request to the `/track` endpoint. It uses the file [long.json](./testdata/long.json). You can change the input file by feeding the variable `DATA` to the command. Example:
```bash
DATA=testdata/short.json make track
```

## Run benchmarks
You can run the benchmarks on your machine (you need Golang installed) with the command: `make bench`.

## Tools
- [trivy](https://github.com/aquasecurity/trivy) was used to identify vulnerabilities in the Dockerfile
- [hadolint](https://github.com/hadolint/hadolint) was used to lint the Dockerfile
- [golangci-lint](https://golangci-lint.run/) was used to lint the Go code

## Potential improvements
The following are ideas that haven't been implemented due to lack of time.

- input data validation: output can be unpredictable if the list of flights is not valid
- config file: it comes handy having the HTTP port configurable. [go-yaml](https://github.com/go-yaml/yaml/tree/v3.0.1) is a good library for `.yaml` config files
- logger: logging has been implemented with the standard library's `log` package. A better alternative would be [logrus](https://github.com/Sirupsen/logrus)
- data generator: the idea was to create a function able to generate a list with hundres/thousands of flights, to better understand which algorithm scales better
- live reload: [reflex](https://github.com/cespare/reflex) is a great tool to test code changes without having to restart the server
- unit tests: self-explanatory
