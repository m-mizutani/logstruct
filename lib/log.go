package logstruct

// Token is a part of log message
type Token struct {
	Data   string `json:"d"`
	freeze bool
}

// Log is a log message
type Log struct {
	Text   string   `json:"text"`
	Tokens []*Token `json:"tokens"`
}

func newLog(line string, t Tokenizer) *Log {
	log := Log{}
	log.Text = line
	log.Tokens = t.Split(line)
	return &log
}

func newToken(d string) *Token {
	t := Token{}
	t.Data = d
	return &t
}

// Equals checks equality
func (x *Token) Equals(t *Token) bool {
	return x.Data == t.Data
}

// Clone creates a copy of original Token
func (x *Token) Clone() *Token {
	c := newToken(x.Data)
	c.freeze = x.freeze
	return c
}

func (x *Token) String() string {
	return x.Data
}
