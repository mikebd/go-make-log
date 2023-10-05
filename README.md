# go-make-log

Generate fake application log files (e.g. as input to log processors).

For safety, this will not overwrite an existing file.

<b>Warning</b>: this does not check if disk space will be exhausted.

## Sample Output

### Create a JSON file (Running in Docker)

```shell
$ docker build -t go-make-log . && docker run -v $(pwd):/output go-make-log -f json -o /output/log.json
```

```shell
$ du -h log.json    # Disk space used
1.0M log.json
```

```shell
$ wc -l log.json    # Number of lines
7531 log.json
```

```shell
$ head -n 3 log.json
{"level":"debug","message":"etisat rif onafon ari y nih ogiwab fosa itecawo to sef weduh baro e halui","timestamp":"2023-10-04T22:03:14.754762Z"}
{"level":"debug","message":"alen ehof h edix ofisew i luwesa ca a hagur get tufiroy wuhoai iseba","timestamp":"2023-10-04T22:03:14.754954Z"}
{"level":"info","message":"p uriye esitos dociha tud oce nosen i sobihas lasobu talena h eto ic","timestamp":"2023-10-04T22:03:14.754994Z"}
```

### Create a 1GiB file (Running Locally)

```shell
$ go run ./main -o log.txt -s 1GiB
```

```shell
$ du -h log.txt     # Disk space used
1.0G log.txt
```

```shell
$ wc -l log.txt     # Number of lines
10578414 log.txt
```

```shell
$ head -n 3 log.txt
2023-10-04T21:46:42.162239Z|info |ehuw ohemo elati l o wogeh arito akeho oheroc potawoh nes ha
2023-10-04T21:46:42.162360Z|debug|godet p zeni rela um beg epaheb bedi ihuyi yor tam el e afuni
2023-10-04T21:46:42.162376Z|debug|erol f o bisow ago r erut sov senas eg ad
```

## Running in Docker

* `docker build -t go-make-log .` - Build the Docker image
* `docker run go-make-log <args>` - Run with the given arguments

## Running Locally

* `go run ./main <args>` - Run the main.go file with the given arguments

## Run Unit Tests

* `go test ./... -count=1` - Run all unit tests

## Arguments

| Argument               | Description                        |
|------------------------|------------------------------------|
| `-h` \| `--help`       | Print the help message             |
| `-lt`                  | Log with timestamps (UTC)          |
| `-f <delimited\|json>` | Output format (default delimited)  |
| `-o <filename>`        | Output filename                    |
| `-s <size>`            | Minimum output size (default 1MiB) |