package encoding

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestRawBinaryEncoder_WriteTime(t *testing.T) {
	assert := asst.New(t)

	var (
		encoder TimeEncoder
		b       []byte
	)
	encoder = NewBigEndianBinaryEncoder()
	num := 100
	b = encodeNanoseconds(t, num, encoder)
	// codec + value
	assert.Equal(8*num+1, len(b))

	encoder = NewLittleEndianBinaryEncoder()
	b = encodeNanoseconds(t, num, encoder)
	assert.Equal(8*num+1, len(b))
}

func TestRawBinaryDecoder_ReadTime(t *testing.T) {
	assert := asst.New(t)
	num := 100
	encoded := encodeNanoseconds(t, num, NewBigEndianBinaryEncoder())
	decoder := NewRawBinaryDecoder()
	assert.Nil(decoder.Init(encoded))
	i := 0
	for decoder.Next() {
		i++
		decoder.ReadTime()
	}
	assert.Equal(i, num)
}

// TODO: use table test and merge with other decoder
func TestRawBinaryDecoder_ReadInt(t *testing.T) {
	assert := asst.New(t)
	encoder := NewBigEndianBinaryEncoder()
	encoder.WriteInt(-1)
	encoder.WriteInt(1)
	decoder := NewRawBinaryDecoder()
	p, err := encoder.Bytes()
	//t.Log(p)
	assert.Nil(err)
	assert.Nil(decoder.Init(p))
	decoder.Next()
	assert.Equal(int64(-1), decoder.ReadInt())
	decoder.Next()
	assert.Equal(int64(1), decoder.ReadInt())
}

// TODO: use table test and merge with other decoder
func TestRawBinaryDecoder_ReadDouble(t *testing.T) {
	assert := asst.New(t)
	encoder := NewBigEndianBinaryEncoder()
	encoder.WriteDouble(-1.1)
	encoder.WriteDouble(1.1)
	decoder := NewRawBinaryDecoder()
	p, err := encoder.Bytes()
	//t.Log(p)
	assert.Nil(err)
	assert.Nil(decoder.Init(p))
	decoder.Next()
	assert.Equal(-1.1, decoder.ReadDouble())
	decoder.Next()
	assert.Equal(1.1, decoder.ReadDouble())
}
