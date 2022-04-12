package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type Ignores []string

var (
	out     io.Writer = os.Stdout
	ignores Ignores
	color   int
	level   int
	dirs    = 0
	files   = 0
	bar     = "├── "
	hor     = "│   "
	space   = "    "
	end     = "└── "
)

func (i *Ignores) String() string {
	return fmt.Sprintf("%+v", *i)
}

func (i *Ignores) Set(v string) error {
	*i = append(*i, v)
	return nil
}

func (i *Ignores) Init() error {
	if len(*i) == 0 {
		*i = []string{".git", "node_modules"}
	}
	return nil
}

func run(dirName string) int {
	if dirName == "" {
		dirName = "."
	}
	//  存在確認
	_, err := os.Stat(dirName)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	fmt.Printf("\x1b[%vm%s\x1b[0m\n", color, dirName)
	if err := tree(0, "", dirName); err != nil {
		fmt.Printf("%+v\n", err)
		return 1
	}
	// ディレクトリ, ファイル総数出力
	fmt.Printf("\n%v directories, %v files\n", dirs, files)
	return 0
}

func tree(l int, prefix, dirName string) error {
	if level != -1 && l >= level {
		return nil
	}
	if slices.Contains(ignores, dirName) {
		return nil
	}
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return errors.Wrap(err, "read dir error:")
	}
	for i, file := range files {
		if file.IsDir() {
			if i == len(files)-1 {
				printDir(prefix+end, filepath.Join(dirName, file.Name()))
				tree(l+1, prefix+space, filepath.Join(dirName, file.Name()))
			} else {
				printDir(prefix+bar, filepath.Join(dirName, file.Name()))
				tree(l+1, prefix+hor, filepath.Join(dirName, file.Name()))
			}
		} else {
			if i == len(files)-1 {
				printFile(prefix+end, filepath.Join(dirName, file.Name()))
				break
			}
			printFile(prefix+bar, filepath.Join(dirName, file.Name()))
		}
	}
	return nil
}

// ファイル出力用
func printFile(mark, fileName string) {
	s := filepath.Base(fileName)
	fmt.Fprintln(out, mark+s)
	files++
}

// ディレクトリ出力用
func printDir(mark, dirName string) {
	s := filepath.Base(dirName)
	fmt.Fprintln(out, mark+fmt.Sprintf("\x1b[%vm%s\x1b[0m", color, s))
	dirs++
}

func init() {
	flag.IntVar(&color, "color", 34, "directory color\n\n30  black\n31  red\n32  green\n33  yellow\n34  blue\n35  magenta\n36  cyaan\n37  white\n")
	flag.IntVar(&level, "L", -1, "level")
	flag.Var(&ignores, "ignore", "ignore keyword")
	flag.Usage = func() {
		fmt.Printf("Usage: %v  [-c color] [-L level] directory \n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	ignores.Init()
	root := flag.Arg(0)
	os.Exit(run(root))
}
