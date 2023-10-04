# go-make-log

Generate fake application log files (e.g. as input to log processors).

For safety, this will not overwrite an existing file.

Warning: this does not check if disk space will be exhausted.

## Running in Docker

* `docker build -t go-make-log .` - Build the Docker image
* `docker run go-make-log <args>` - Run with the given arguments

## Running Locally

* `go run ./main <args>` - Run the main.go file with the given arguments

## Arguments

| Argument                | Description                       |
|-------------------------|-----------------------------------|
| `-h` \| `--help`        | Print the help message            |
| `-lt`                   | Log with timestamps (UTC)         |
| `-f <delimited\|json>`  | Output format (default delimited) |
| `-o filename`           | Output filename                   |