# Monit

Monit is a small tool created to allows you to do really basic monitoring using bash commands.
It allows you to schedule command execution and do comparaison against returned values of thoses bash commands.
You can then, make call on its API to get monitored commands status.

## Configuration

You should pass a path to a YAML file when launching the process.

```sh
monit -config ./config.yaml
```

The YAML file should look like this:

```yaml
http:
  listenAddr: '0.0.0.0:8000'
  workingResponse: 'OK'
  errorResponse: 'KO'
services:
  - name: DISK
    checkJob:
      command: "df -Ph . | tail -1 | awk '{print $5}' | cut -c1-2"
      cron: '* * * * * *'
      test: 'gt'
      value: '80'
  - name: RAM
    checkJob:
      command: 'echo 60'
      cron: '* * * * * *'
      test: 'eq'
      value: '60'
```

_Parameters:_

| name                         | Description                                                                        |
| ---------------------------- | ---------------------------------------------------------------------------------- |
| http.listenAddr              | The listen address for API and metrics                                             |
| http.workingResponse         | The text to return if the comparaison against the command's return value was true  |
| http.errorResponse           | The text to return if the comparaison against the command's return value was false |
| services.[].name             | The name of the service (for API routes)                                           |
| services.[].checkJob.command | The command to run                                                                 |
| services.[].checkJob.cron    | The cron pattern (you can use format of https://godoc.org/gopkg.in/robfig/cron.v2) |
| services.[].checkJob.test    | The comparaison to run                                                             |
| services.[].checkJob.value   | The compared value for your test                                                   |

## API

For each of the services, the program exposes the following route `/services/[NAME]`.
Each route returns http.workingResponse if the test for your service worked other it returns http.errorResponse
