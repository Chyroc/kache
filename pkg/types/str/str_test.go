package str

import (
	"testing"

	"bytes"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {
	assert := testifyAssert.New(t)

	{
		// New
		str := New()

		// Append
		str.Append('a')
		str.Append('b')
		str.Append('1')
		str.Append('2')

		assert.Equal("ab12", str.String())
		assert.False(str.Empty())
		assert.Equal(4, str.Len())

		// Reset
		str.Reset()

		assert.Equal("", str.String())
		assert.True(str.Empty())
		assert.Equal(0, str.Len())

		// AppendString
		str.AppendString("hello")

		assert.Equal("hello", str.String())
		assert.Equal([]byte("hello"), str.Bytes())
		assert.False(str.Empty())
		assert.Equal(5, str.Len())

		// NewFromString
		str = NewFromString("hello")

		str.Append(' ')
		str.AppendString(`world`)
		str.AppendString(` hha `)

		assert.Equal("hello world hha ", str.String())
		assert.Equal([]byte("hello world hha "), str.Bytes())
		assert.False(str.Empty())
		assert.Equal(16, str.Len())
	}

	// long string
	{
		s := ""
		for i := 0; i < 100; i++ {
			s += "test string"
		}

		str := NewFromString(s)
		assert.False(str.Empty())
		assert.Equal(1100, str.len)
		assert.Equal(1100, str.cap)

		// append 2200 string to 1100 cap Str, more than 2 times, to expect length
		str.AppendString(s + s)

		assert.Equal(3300, str.len)
		assert.Equal(3300, str.cap)

		// append 1 string to 3300 cap Str, less than 2 times, double
		str.AppendString(" ")

		assert.Equal(3301, str.len)
		assert.Equal(6600, str.cap)
	}
}

func BenchmarkStr_AppendString(b *testing.B) {
	s := "this is for test string"
	appendTimes := 20

	b.Run("pure string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str string
			for j := 0; j < appendTimes; j++ {
				str += s
			}
		}
	})

	b.Run("bytes buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str = new(bytes.Buffer)
			for j := 0; j < appendTimes; j++ {
				str.WriteString(s)
			}
			_ = str.String()
		}
	})

	b.Run("optimized string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str = New()
			for j := 0; j < appendTimes; j++ {
				str.AppendString(s)
			}
			_ = str.String()
		}
	})
}

func BenchmarkStr_AppendChar(b *testing.B) {
	appendTimes := 20

	b.Run("pure string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str string
			for j := 0; j < appendTimes; j++ {
				str += " "
			}
		}
	})

	b.Run("bytes buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str = new(bytes.Buffer)
			for j := 0; j < appendTimes; j++ {
				str.WriteByte(' ')
			}
			_ = str.String()
		}
	})

	b.Run("optimized string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str = New()
			for j := 0; j < appendTimes; j++ {
				str.Append(' ')
			}
			_ = str.String()
		}
	})
}
