# tree
devquiz in [golang.tokyo#21](https://techplay.jp/event/712820)

# usage
```
Usage: mytree  [-c color] [-L level] directory
  -L int
        level
  -color int
        directory color

        30  black
        31  red
        32  green
        33  yellow
        34  blue
        35  magenta
        36  cyaan
        37  white
         (default 34)
```

# run
```
$  go run ./cmd/mytree [-c color] [-L level] directory
```

# test
```
$ go test ./cmd/mytree
```

# build
```
$ go build -o bin/mytree ./cmd/mytree 
```
