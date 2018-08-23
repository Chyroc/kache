package str

const (
	DefaultStrLength = 32
)

type Str struct {
	content []byte
	len     int
	cap     int
}

// New return a Str interface
func New() *Str {
	return &Str{content: make([]byte, DefaultStrLength), cap: DefaultStrLength}
}

// NewFromString return a Str interface by a string
func NewFromString(s string) *Str {
	sLen := len(s)
	if sLen > 32 {
		return &Str{content: []byte(s), len: sLen, cap: sLen}
	}

	content := make([]byte, DefaultStrLength)
	copy(content[:sLen], s)
	return &Str{content: content, len: sLen, cap: DefaultStrLength}
}

func (s *Str) Append(b byte) {
	if s.len+1 > s.cap {
		s.cap *= 2
		content := make([]byte, s.cap*2)
		copy(content[:s.len], s.content)
		s.content = content
	}

	s.content[s.len] = b
	s.len++
}

func (s *Str) AppendString(str string) {
	var strLen = len(str)
	var expectLen = s.len + strLen

	if expectLen > s.cap {
		c := s.cap * 2
		if c < expectLen {
			c = expectLen
		}
		s.cap = c

		content := make([]byte, s.cap)
		copy(content[:s.len], s.content)
		s.content = content
	}

	copy(s.content[s.len:expectLen], []byte(str)[:])
	s.len = expectLen
}

func (s *Str) String() string {
	return string(s.content[:s.len])
}

func (s *Str) Bytes() []byte {
	return s.content[:s.len]
}

func (s *Str) Len() int {
	return s.len
}

func (s *Str) Empty() bool {
	return s.len == 0
}

func (s *Str) Reset() {
	s.len = 0
	s.content = make([]byte, DefaultStrLength)
	s.cap = DefaultStrLength
}
