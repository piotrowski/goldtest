# goldtest

[![GoDoc](https://godoc.org/github.com/xorcare/golden?status.svg)](https://godoc.org/github.com/apiotrowski312/goldtest)
[![Go Report Card](https://goreportcard.com/badge/github.com/apiotrowski312/goldtest)](https://goreportcard.com/report/github.com/apiotrowski312/goldtest)
[![codecov](https://codecov.io/gh/xorcare/golden/badge.svg)](https://codecov.io/gh/apiotrowski312/goldtest)
![coverage](https://github.com/apiotrowski312/goldtest/workflows/Build/badge.svg)

goldtest is a simple package that improves experience of working with golden files. It allows you to assert any type of data with saved golden file. Currently there are three exported functions, I believe it should be enough for most cases. If you see room for improvement, please leave an issue or pull request.

## What is Golden file?

Golden file is another way to store "expected output" of a function. One "golden file" stands for output from one test case. Usually those files are stored in separated files, in directory "testdata".

Example usage is in [example](example/) directory. Also you can read a little bit more about package in [THIS](https://apiotrowski.xyz/posts/packages/goldtest/) blog post

You can read more about it here:

- https://medium.com/@jarifibrahim/golden-files-why-you-should-use-them-47087ec994bf
- https://softwareengineering.stackexchange.com/questions/358786/what-is-golden-files

## Usage

You can update(create) golden files automatically by providing flag `-update`.

There are two main functions:

```
Assert() - good for smaller files, there is no need to have any field exported, works good with []byte.
AssertJSON() - works good for big structs (E.g tree), will give you more readible golden file. If you want to use it, all fields that should be tested should be exported.
```

Both works similar for build-in types (Eg. string, int, float32 etc.). Output is a bit different for maps, structs and bytes.

I suggest using JSON when you have to work with big data structures as it will be more readable.

> **Note**: As you can see here: [Assert](testdir/testfile_byte.golden) [AssertJSON](testdir/json_testfile_byte.golden), there is wrong output, when using AssertJSON(). This is how it is, so if you want use golden_file for bytes, consider using Assert()

Example can be found in [example directory](example/main_test.go)

