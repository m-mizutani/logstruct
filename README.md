# logstruct [![Travis-CI](https://travis-ci.org/m-mizutani/logstruct.svg)](https://travis-ci.org/m-mizutani/logstruct) [![Report card](https://goreportcard.com/badge/github.com/m-mizutani/logstruct)](https://goreportcard.com/report/github.com/m-mizutani/logstruct) 

## What is this

`logstruct` parses and estimates log original format like a `printf` format argument from existing text log data.

## Getting Started

At first, install logstrut with following command.

```bash
get github.com/m-mizutani/logstruct
```

Invoke `logstruct` command and output formats.

```bash
$ logstruct /var/log/auth.log
0 : Dec * *:*:* pylon sshd[*]: Invalid user  from 139.162.122.110
1 : Dec * *:*:* pylon sshd[*]: Invalid user * from *
2 : Dec * *:*:* pylon sshd[*]: Connection closed by * [preauth]
3 : Dec * *:*:* pylon sshd[*]: input_userauth_request: invalid user  [preauth]
4 : Dec * *:*:* pylon sshd[*]: input_userauth_request: invalid user * [preauth]
5 : Dec 27 19:*:* pylon sshd[*]: fatal: no hostkey alg [preauth]
(snip)
```

### Export and import model (formats and metadata)

Export model and save it to a file.

```bash
logstruct -e exported.model  /var/log/auth.log
```

And import it.

```bash
logstruct -i exported.model  /var/log/auth.2.log
```

### Change log level

`--log-level`  or `-l` option can choose log leve from `debug`, `info` and `warn`.

```bash
logstruct -l info /var/log/auth.log
```