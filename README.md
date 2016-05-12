# cwlogs-tailf

Follow tail messages of CloudWatch Logs.

## Usage

```
Usage of cwlogs-tailf:
  -g string
      log group name
  -s string
      log stream name
```

```sh
$ cwlogs-tailf -g my-group -s my-stream
2016-05-13T01:51:18+09:00 foo
2016-05-13T01:51:18+09:00 bar
2016-05-13T01:51:19+09:00 zoo
...
```

## Installation

```
brew install https://raw.githubusercontent.com/winebarrel/cwlogs-tailf/master/homebrew/cwlogs-tailf.rb
```
