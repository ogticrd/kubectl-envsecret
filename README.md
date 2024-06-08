# kubectl-envsecret

## Overview

`kubectl-envsecret` is a Kubernetes plugin designed to simplify the creation of
Kubernetes secrets from `.env` files, including support for multiline
environment variables. This tool streamlines the management of secrets, making
it easier to handle configurations that include complex, multiline values.

## Features

- Create Kubernetes secrets from `.env` files.
- Support for multiline environment variables.
- Easy integration with `kubectl` as a plugin.

## Installation

To install `kubectl-envsecret`, you need to have Go installed on your machine.
You can then use the following command to install the plugin:

```sh
go install github.com/ogticrd/kubectl-envsecret@latest
```

## Usage

### Basic Usage

The primary command provided by `kubectl-envsecret` is `create`, which reads a
`.env` file and creates a corresponding Kubernetes secret.

```sh
kubectl envsecret create --from-env-file <path-to-env-file> --from-env-file <path-to-other-files>
```

### Command Options

- `--from-env-file`: Specifies the path(s) to the `.env` file. This option can
  be used multiple times to specify multiple `.env` files.

### Examples

#### Create a Secret from a Single `.env` File

```sh
kubectl envsecret create --from-env-file /path/to/`.env`
```

#### Create a Secret from Multiple `.env` Files

```sh
kubectl envsecret create --from-env-file /path/to/`.env` --from-env-file /another/path/.env
```

## Development

### Prerequisites

- Go
- Cobra CLI

### Downloading the project

First you need is cloning the project

```sh
git clone https://github.com/ogticrd/kubectl-envsecret.git
cd kubectl-envsecret
```

### Building the Project

To build the project, clone the repository and run:

#### With `go` command

```sh
go build -o kubectl-envsecret main.go
```

#### Using `make` command

```sh
make build
```

### Running the Project

#### With `go` command

```sh
go run main.go
```

> [!NOTE]\
> Just run `./kubectl-envsecret` if you already built it.

#### Using `make` command

```sh
make run
```

### Running Tests

To run the tests, use the following command:

#### With `go` command

```sh
go test ./...
```

#### Using `make` command

```sh
make test
```

## Project Structure

```plaintext
.
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── create.go
│   └── root.go
├── go.mod
├── go.sum
├── internal
│   ├── k8sapi
│   │   └── k8s.go
│   ├── parser
│   │   └── env.go
│   └── utils
│       └── utils.go
└── main.go
```

- **cmd**: Contains the CLI command definitions.
- **internal/k8sapi**: Contains a wrapper of the usage of Kubernetes API to
  manage secrets.
- **internal/parser**: Contains functions to parse `.env` files.
- **internal/utils**: Contains utility functions used by the commands.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file
for details.

## Acknowledgements

This project uses the following libraries:

- [Cobra CLI](https://github.com/spf13/cobra) for creating the command-line
  interface.
- [Kubernetes CLI Runtime](https://github.com/kubernetes/cli-runtime) for
  Kubernetes API interactions.
