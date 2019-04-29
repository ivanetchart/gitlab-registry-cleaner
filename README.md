# Gitlab Registry Cleaner

A simple tool intended to keep gitlab registries in shape.
Currently Gitlab is not enforcing the 10gb repository size limit, so you can store
any images as you want. In order to make sure this doesn't happen and eventually there won't be any
issue you can use this tool to keep your registries clean.

## Installation

```
$ go get github.com/ivanetchart/gitlab-registry-cleaner
```

Assuming you have your `GOPATH` setup already (as part of your `PATH` env var), you will able to run it without needing to directly go to `$GOPATH/bin`.

## Usage

Default (basic) usage:

```
$ gitlab-registry-cleaner -project=<your-gitlab-project> -token=<a-gitlab-user-token>
```

For more available configurations:

```
$ gitlab-registry-cleaner --help
```








