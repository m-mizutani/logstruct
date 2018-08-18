package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type options struct {
	// MaxLen uint   `long:"maxlen" description:"Max length of log message"`
	Output      string  `short:"o" long:"output" description:"Output file, '-' means stdout" default:"-"`
	Dumper      string  `short:"d" long:"dumper" choice:"text" choice:"json" choice:"sjson" choice:"heatmap" default:"text"`
	Threshold   float64 `short:"t" long:"threshold" default:"0.7"`
	Delimiters  string  `short:"s" long:"delimiters"`
	Content     string  `short:"c" long:"content" choice:"log" choice:"format" default:"format"`
	EnableRegex bool    `long:"enable-regex"`
	// FileName string `short:"i" description:"A log file" value-name:"FILE"`
}

func main() {
	var opts options

	args, ParseErr := flags.ParseArgs(&opts, os.Args)
	if ParseErr != nil {
		os.Exit(1)
	}

	fmt.Println(args)

}
