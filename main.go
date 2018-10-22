package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	logstruct "github.com/m-mizutani/logstruct/lib"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type options struct {
	// MaxLen uint   `long:"maxlen" description:"Max length of log message"`
	// Output      string  `short:"o" long:"output" description:"Output file, '-' means stdout" default:"-"`
	Threshold  float64 `short:"t" long:"threshold" default:"0.7"`
	ImportFile string  `short:"i" long:"import"`
	ExportFile string  `short:"e" long:"export"`
	Dump       bool    `shosrt:"d" long:"dump-format"`
}

func readFile(m *logstruct.Model, fpath string) error {
	fd, err := os.Open(fpath)
	if err != nil {
		return errors.Wrap(err, "Fail to open log file: "+fpath)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		format, isNew := m.InputLog(scanner.Text())
		if isNew {
			log.WithField("log", format.String()).Info("New Format")
		}
	}

	return nil
}

type logstructOpt struct {
	modelInput  string
	modelOutput string
	files       []string
}

func main() {
	app := cli.NewApp()
	app.Name = "logstruct"
	app.Usage = "Automatci log format estimator"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Input model",
		},
	}
	app.Action = func(c *cli.Context) error {
		args := []string{}
		for i := 0; i < c.NArg(); i++ {
			args = append(args, c.Args().Get(i))
		}

		opts := logstructOpt{
			modelInput:  c.String("i"),
			modelOutput: c.String("e"),
			files:       args,
		}

		return logstructMain(opts)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func logstructMain(opts logstructOpt) error {
	m := logstruct.NewModel()

	if opts.modelInput != "" {
		err := m.Load(opts.modelInput)
		if err != nil {
			return err
		}
	}

	if opts.modelOutput != "" {
		defer func() {
			err := m.Save(opts.modelOutput)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	for _, arg := range opts.files {
		err := readFile(m, arg)
		if err != nil {
			return err
		}
	}

	for idx, format := range m.Formats() {
		fmt.Println(idx, ":", format)
	}

	return nil
}
