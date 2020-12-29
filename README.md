# Go Docker Operator

Docker operator written in Go

## Usage

Need declare these environment variable:

```
DOCKER_HOST=https://hub.docker.com
DOCKER_API_VERSION=1.41
```

### Run by file

```shell
go mod download
go build main.go
DOCKER_HOST=tcp://host.docker.internal:2375 DOCKER_API_VERSION=1.41 ./main <COMMAND> [OPTIONS]
```

### Run by Docker image

```shell
docker run --rm -e DOCKER_HOST=tcp://host.docker.internal:2375 -e DOCKER_API_VERSION=1.41 gocker-operator:0.0.1 prune <COMMAND> [OPTIONS]
```

## Commands

### Prune

Prune images

|   Options   | Default value |                     Description                      |
| ----------- | ------------- | ---------------------------------------------------- |
| `-dangling` | `true`        | `true`: remove dangling images, `false`: remove both |
| `-until`    | `30`          | Number of days to keep                               |