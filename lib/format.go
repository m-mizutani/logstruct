package logstruct

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// Format has log data structure information
type Format struct {
	BaseLog *Log
	Blocks  []*Block
}

// Block is a part of format template
type Block struct {
	Data    string `json:"d"`
	IsParam bool   `json:"p"`
}

func (x *Block) String() string {
	if x.IsParam {
		return "*"
	}

	return x.Data
}

// NewFormat is a constructor of Format
func NewFormat(baseLog *Log) *Format {
	f := Format{}
	f.BaseLog = baseLog
	f.Blocks = make([]*Block, len(baseLog.Tokens))
	for idx, token := range baseLog.Tokens {
		f.Blocks[idx] = &Block{token.Data, false}
	}

	return &f
}

// Match returns matching ratio of a log and the format
func (x *Format) Match(log *Log) float64 {
	matched := 0

	if len(x.BaseLog.Tokens) != len(log.Tokens) || len(x.BaseLog.Tokens) != len(x.Blocks) {
		logrus.WithFields(logrus.Fields{
			"x.BaseLog.Tokens": len(x.BaseLog.Tokens),
			"log.Tokens":       len(log.Tokens),
			"x.Blockes":        len(x.Blocks),
		}).Fatal("Length mismatch")
	}

	for idx, token := range log.Tokens {
		if token.Data == x.Blocks[idx].Data {
			matched++
		}
	}

	return float64(matched) / float64(len(x.BaseLog.Tokens))
}

func (x *Format) Merge(log *Log) {
	before := x.String()
	changed := false
	for idx, token := range log.Tokens {
		if token.Data != x.Blocks[idx].Data && !x.Blocks[idx].IsParam {
			logrus.WithFields(logrus.Fields{
				"t1": x.Blocks[idx].Data,
				"t2": token.Data,
			}).Debug("Changed token")
			x.Blocks[idx].IsParam = true
			changed = true
		}
	}

	if changed {
		logrus.WithFields(logrus.Fields{
			"base":   x.BaseLog.Text,
			"before": before,
			"after":  x.String(),
		}).Debug("Changed")
	}
}

func (x *Format) String() string {
	arr := make([]string, len(x.Blocks))
	for idx, block := range x.Blocks {
		arr[idx] = block.String()
	}

	return strings.Join(arr, "")
}

/*
func (x *Format) ColoredString() string {
}
*/
