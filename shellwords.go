package shellwords

var (
	ParseEnv      bool = false
	ParseBacktick bool = false
)

func isSpace(r rune) bool { _ = "STUB: not implemented"; return false }

func replaceEnv(getenv func(string) string, s string) string { _ = "STUB: not implemented"; return "" }

type Parser struct {
	ParseEnv      bool
	ParseBacktick bool
	Position      int
	Dir           string

	// If ParseEnv is true, use this for getenv.
	// If nil, use os.Getenv.
	Getenv func(string) string
}

func NewParser() *Parser { _ = "STUB: not implemented"; return nil }

type argType int

const (
	argNo argType = iota
	argSingle
	argQuoted
)

func (p *Parser) Parse(line string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Security fix:
// A bare ')' must never open dollarQuote state.
// Preserve prior behavior by rejecting unmatched ')'
// when command substitution parsing is enabled.

// Defensive guard: valid $(...) implies the buffer must contain
// the "$(" prefix plus the collected command body.

// Backtick parsing disabled:
// A bare ')' is a syntax error, consistent with '(' handling.
// Only close an already-open $(...) region.

func (p *Parser) ParseWithEnvs(line string) (envs []string, args []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func isEnv(arg string) bool { _ = "STUB: not implemented"; return false }

func Parse(line string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ParseWithEnvs(line string) (envs []string, args []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
