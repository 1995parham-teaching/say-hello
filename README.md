# say-hello

## Introduction

Sometimes we need to have a frontend besides the backend.
This repository uses React to show how we can server a React build and then use API for communication.

Also we are using the [go 1.16 embed package](https://pkg.go.dev/embed) so you can use the binary with frontend
where ever you want.

```sh
# starting from the project root directory
cd web/say-hello
npm build

cd ../..
go build
```
