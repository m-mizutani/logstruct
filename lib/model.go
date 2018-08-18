package logstruct

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

// Model is a main structure of logstruct
type Model struct {
	tokenizer Tokenizer
	FormatMap map[int][]*Format `json:"fmap"`
	Threshold float64           `json:"threshold"`
}

// NewModel is a constructor of Model
func NewModel() *Model {
	m := Model{}
	m.FormatMap = make(map[int][]*Format)
	m.tokenizer = NewSimpleTokenizer()
	m.Threshold = 0.65
	return &m
}

// InputLog reads log message and import to model.
func (x *Model) InputLog(msg string) (matchedForamt *Format, newFormat bool) {
	log := newLog(msg, x.tokenizer)

	maxRatio := 0.0
	length := len(log.Tokens)

	for _, format := range x.FormatMap[length] {
		ratio := format.Match(log)
		if ratio > maxRatio {
			matchedForamt = format
		}
	}

	if matchedForamt == nil {
		matchedForamt = NewFormat(log)
		newFormat = true
		if x.FormatMap[length] == nil {
			x.FormatMap[length] = make([]*Format, 0)
		}
		x.FormatMap[length] = append(x.FormatMap[length], matchedForamt)
	}

	return
}

// Export stores to byte array
func (x *Model) Export() ([]byte, error) {
	data, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Import restore from byte array
func (x *Model) Import(data []byte) error {
	err := json.Unmarshal(data, &x)
	return err
}

// Save stores model to a file
func (x *Model) Save(fpath string) error {
	data, err := x.Export()
	if err != nil {
		return err
	}

	fd, err := os.Create(fpath)
	if err != nil {
		return errors.Wrap(err, "Fail to create a model file")
	}

	n, err := fd.Write(data)
	if err != nil {
		return errors.Wrap(err, "Fail to write data to a model file")
	}

	if n != len(data) {
		return errors.New("Invalid write size")
	}

	return nil
}

// Load restores model from a file
func (x *Model) Load(fpath string) (err error) {
	return
}

// Formats returns all format of the model
func (x *Model) Formats() []*Format {
	arr := make([]*Format, 0)
	for _, formats := range x.FormatMap {
		arr = append(arr, formats...)
	}
	return arr
}
