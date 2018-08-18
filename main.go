package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	logstruct "github.com/m-mizutani/logstruct/lib"
	"github.com/pkg/errors"
)

type options struct {
	// MaxLen uint   `long:"maxlen" description:"Max length of log message"`
	// Output      string  `short:"o" long:"output" description:"Output file, '-' means stdout" default:"-"`
	Threshold  float64 `short:"t" long:"threshold" default:"0.7"`
	ImportFile string  `short:"i" long:"import"`
	ExportFile string  `short:"e" long:"export"`
}

func readFile(m *logstruct.Model, fpath string) error {
	fd, err := os.Open(fpath)
	if err != nil {
		return errors.Wrap(err, "Fail to open log file: "+fpath)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		m.InputLog(scanner.Text())
	}

	return nil
}

func main() {
	var opts options

	args, ParseErr := flags.ParseArgs(&opts, os.Args)
	if ParseErr != nil {
		os.Exit(1)
	}

	m := logstruct.NewModel()

	if opts.ImportFile != "" {
		err := m.Load(opts.ImportFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	if opts.ExportFile != "" {
		defer func() {
			err := m.Save(opts.ExportFile)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	for _, arg := range args[1:] {
		err := readFile(m, arg)
		if err != nil {
			log.Fatal(err)
		}
	}

	for idx, format := range m.Formats() {
		fmt.Println(idx, ":", format)
	}

}
