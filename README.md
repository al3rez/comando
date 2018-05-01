# comando

💂 Command as a service

## Install

```
~ $ go get github.com/cooldrip/comando/cmd/comando
```

## CLI Example

Pass a command to `comando` with a endpoint:

```
~ $ comando -c "echo foo" -r foo
```

Now you can go to `localhost:6000/foo` and get the output:

```json
{ "output": "foo" }
```

Outputs are trimmed by default with `--raw-output` outputs will be returned they were!

```
~ $ comando -c "echo foo" -r foo --raw-output
```

```json
{ "output": "foo\n" }
```

## CLI Usage

```
NAME:
   comando - Command as a service

USAGE:
   comando [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c value       Command to execute (default: "echo 'as you were!'")
   -r value       Route to serve on the command output (default: "comando")
   --raw-output   Output will be returned directly as a JSON string
                  rather than being trimmed.
   --port value   Set port (default: "6000")
   --host value   Set host (default: "localhost")
   --help, -h     show help
   --version, -v  print the version
```
