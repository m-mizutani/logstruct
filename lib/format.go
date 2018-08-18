package logstruct

// Format has log data structure information
type Format struct {
	baseLog *Log
}

// NewFormat is a constructor of Format
func NewFormat(baseLog *Log) *Format {
	f := Format{}
	f.baseLog = baseLog
	return &f
}

// Match returns matching ratio of a log and the format
func (x *Format) Match(log *Log) float64 {
	matched := 0
	for idx, token := range x.baseLog.Tokens {
		if token.Equals(log.Tokens[idx]) {
			matched++
		}
	}

	return float64(matched) / float64(len(x.baseLog.Tokens))
}
