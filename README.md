# Say Hello :wave

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/1995parham-teaching/say-hello/test.yaml?label=test&logo=github&style=flat-square&branch=main)
[![Codecov](https://img.shields.io/codecov/c/gh/1995parham-teaching/say-hello?logo=codecov&style=flat-square)](https://codecov.io/gh/cng-by-example/say-hello)

## Introduction

Sometimes we need to have a frontend besides the backend.
This repository uses React to show how we can server a React build and then use API for communication.

Also, we are using the [go 1.16 embed package](https://pkg.go.dev/embed), so you can use the binary with frontend
where ever you want.

```bash
# starting from the project root directory
cd web/say-hello
npm run build

cd ../..
go build
```
