# cwlogs-tailf

Follow tail messages of CloudWatch Logs.

## Usage

```
Usage of cwlogs-tailf:
  -V verbose output
  -g string
      log group name
  -f string
      log stream name filter regexp
  -s string
      log stream name
```

```sh
$ cwlogs-tailf -V -g my-group -s my-stream
2016-05-13T01:51:18+09:00 foo
2016-05-13T01:51:18+09:00 bar
2016-05-13T01:51:19+09:00 zoo
...
```

```sh
$ cwlogs-tailf -V -g my-group -f ^stream-prefix-
stream1 2016-05-13T01:51:18+09:00 foo
stream2 2016-05-13T01:51:18+09:00 bar
stream1 2016-05-13T01:51:19+09:00 zoo
...
```

## Installation

```
brew install https://raw.githubusercontent.com/winebarrel/cwlogs-tailf/master/homebrew/cwlogs-tailf.rb
```
