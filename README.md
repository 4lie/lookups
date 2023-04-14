# Geo Lookups

[![Build](https://github.com/4lie/lookups/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/4lie/lookups/actions/workflows/build.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/4lie/lookups)](https://pkg.go.dev/github.com/4lie/lookups)

## Introduction

Fast and in-memory geo lookup library.

Simply add polygons and run queries.
It uses Google's S2 Library for indexing and it's super fast :rocket:.

## Install

``` bash
go get github.com/4lie/lookups
```

## Examples

`lookups_test.go` contains several useful examples. You can check visualizations of them here.

| Image                                       | Description                |
|---------------------------------------------|----------------------------|
| <img src="./images/index.png" width="350"/> | How indexing works         |
| <img src="./images/1.png" width="350"/>     | Simple polygon example     |
| <img src="./images/2.png" width="350"/>     | Polygon with holes example |
| <img src="./images/3.png" width="350"/>     | Multiple polygons example  |
