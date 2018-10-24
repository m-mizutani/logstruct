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
	verbose     bool
	threshold   float64
	files       []string
	quiet       bool
}

func main() {
	app := cli.NewApp()
	app.Name = "logstruct"
	app.Usage = "Automatci log format estimator"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "import, i",
			Usage: "A file imported model",
		},
		cli.StringFlag{
			Name:  "export, e",
			Usage: "A file exported model",
		},
		cli.StringFlag{
			Name:  "log-level, l",
			Usage: "Set log level (warn, info, debug)",
			Value: "warn",
		},
		cli.Float64Flag{
			Name:  "threshold, t",
			Usage: "Threshold to create a new format cluster",
		},
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "disable output formats at last",
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
			threshold:   c.Float64("t"),
			quiet:       c.Bool("q"),
		}

		switch c.String("l") {
		case "warn":
			log.SetLevel(log.WarnLevel)
		case "info":
			log.SetLevel(log.InfoLevel)
		case "debug":
			log.SetLevel(log.DebugLevel)
		default:
			log.WithField("given log level", c.String("l")).Fatal("Invalid log level")
		}

		return logstructMain(opts)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func logstructMain(opts logstructOpt) error {
	if opts.verbose {
		log.SetLevel(log.DebugLevel)
	}

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

	if opts.threshold > 0 {
		m.Threshold = opts.threshold
	}

	log.WithField("threshold", m.Threshold).Info("Model Parameters")
	for _, arg := range opts.files {
		err := readFile(m, arg)
		if err != nil {
			return err
		}
	}

	if !opts.quiet {
		for idx, format := range m.Formats() {
			fmt.Println(idx, ":", format)
		}
	}

	return nil
}
