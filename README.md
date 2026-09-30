# simple-log

Formats JSON log lines; other lines pass through unchanged.

Run a command and format its stdout and stderr separately:

```sh
simple-log -- ./app arg1 arg2
```

The command receives standard input. `simple-log` forwards SIGINT and SIGTERM and
exits with the command's status. If the command is terminated by a signal, the
exit status is `128 + signal number`.

The original stdin mode is also available:

```sh
some-command | simple-log
```

For a Docker image containing both binaries, the entrypoint can be:

```dockerfile
ENTRYPOINT ["simple-log", "--", "./app"]
```
