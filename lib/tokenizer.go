package logstruct

import (
	// "log"
	"regexp"
	"strings"
)

// Tokenizer splits log message string
type Tokenizer interface {
	Split(msg string) []*Token
}

// SimpleTokenizer is one of implementation for Tokenizer
type SimpleTokenizer struct {
	delims    string
	regexList []*regexp.Regexp
	useRegex  bool
}

// NewTokenizer is a wrapper of SimpleTokenizer
func NewTokenizer() Tokenizer {
	return NewSimpleTokenizer()
}

// NewSimpleTokenizer is a constructor of SimpleTokenizer
func NewSimpleTokenizer() *SimpleTokenizer {
	s := &SimpleTokenizer{}
	s.delims = " \t!,:;[]{}()<>=|\\*\"'"
	s.useRegex = true

	heuristicsPatterns := []string{
		// DateTime
		`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.\d+`,
		`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`,
		`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`,
		// Date
		`\d{4}/\d{2}/\d{2}`,
		`\d{4}-\d{2}-\d{2}`,
		`\d{2}:\d{2}:\d{2}.\d+`,
		// Time
		`\d{2}:\d{2}:\d{2}`,
		// Mail address
		`[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*`,
		// IPv4 address
		`(\d{1,3}\.){3}\d{1,3}`,
	}

	s.regexList = make([]*regexp.Regexp, len(heuristicsPatterns))
	for idx, p := range heuristicsPatterns {
		s.regexList[idx] = regexp.MustCompile(p)
	}
	return s
}

// SetDelim is a function set characters as delimiter
func (x *SimpleTokenizer) SetDelim(d string) {
	x.delims = d
}

// EnableRegex is disabler of heuristics patterns
func (x *SimpleTokenizer) EnableRegex() {
	x.useRegex = true
}

// DisableRegex is disabler of heuristics patterns
func (x *SimpleTokenizer) DisableRegex() {
	x.useRegex = false
}

func (x *SimpleTokenizer) splitByDelimiter(chunk *Token) []*Token {
	var res []*Token
	msg := chunk.Data
	for {
		idx := strings.IndexAny(msg, x.delims)
		if idx < 0 {
			if len(msg) > 0 {
				res = append(res, newToken(msg))
			}
			break
		}

		// log.Print("index: ", idx)
		fwd := idx + 1

		s1 := msg[:idx]
		s2 := msg[idx:fwd]
		s3 := msg[fwd:]

		if len(s1) > 0 {
			// log.Print("add s1: ", s1)
			res = append(res, newToken(s1))
		}

		if len(s2) > 0 {
			// log.Print("add s2: ", s2)
			res = append(res, newToken(s2))
		}

		msg = s3
		// log.Print("remain: ", msg)
	}

	return res
}

// Split is a function to split log message.
func (x *SimpleTokenizer) Split(msg string) []*Token {
	token := newToken(msg)
	chunks := []*Token{token}

	var res []*Token
	for _, c := range chunks {
		if c.freeze {
			res = append(res, c)
		} else {
			res = append(res, x.splitByDelimiter(c)...)
		}
	}
	return res
}
