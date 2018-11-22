# Log - a pretty basic logger.

A Go logger. Very basic. There are probably better loggers out there. Nothing much else to see here.

## What?
A basic logger which writes a variable number of delimeted strings to stderr. Used by [auth].

## Why?
Provides logging to stderr.

## How?
See [examples].

## Examples
See [examples] for a http/appengine implementations which uses log and auth. This is written for appengine standard 2nd gen, but also works as a standalone.

## Dependencies and services
None.

## Installation
If you want to run the example code, then install using
```sh
$ go get -u github.com/lidstromberg/examples
```
If you only want the log utility, then install with
```sh
$ go get -u github.com/lidstromberg/log
```


   [auth]: <https://github.com/lidstromberg/auth>
   [examples]: <https://github.com/lidstromberg/examples>