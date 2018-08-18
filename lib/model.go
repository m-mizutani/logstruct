package logstruct

import (
	"encoding/json"
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
func (x *Model) Save(fpath string) (err error) {
	return
}

// Load restores model from a file
func (x *Model) Load(fpath string) (err error) {
	return
}
