package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T) {

	value := "2"

	err := Set(strconv.FormatInt(380670000000, 10), "billing", value)
	assert.Nil(t, err)

	billing, err := Get(strconv.FormatInt(380670000000, 10), "billing")
	assert.NotNil(t, billing)
	assert.Nil(t, err)

	assert.EqualValues(t, string(billing), value)
}

// mac noatime
// BenchmarkReadOpen-8		300_000_000		3.58 ns/op
// BenchmarkWriteOpen-8  	    500_000   2919 ns/op
func BenchmarkRead(b *testing.B) {
	b.ResetTimer()

	from := 380670000000

	for id := int64(from); id < int64(b.N+from); id++ {
		if _, err := Get(strconv.FormatInt(id, 10), "billing"); err != nil {
			log.Println(err)
		}
	}
}

func BenchmarkWrite(b *testing.B) {
	b.ResetTimer()

	from := 380670000000

	for id := int64(from); id < int64(100000+from); id++ {
		if err := Set(strconv.FormatInt(id, 10), "lang", strconv.FormatInt(rand.Int63n(3), 10)); err != nil {
			log.Println(err)
		}
	}
}

func BenchmarkReadOpen(b *testing.B) {

	if f, err := os.OpenFile("id", os.O_RDONLY|os.O_CREATE, 600); err != nil {
		log.Fatal(err)
	} else {

		b.ResetTimer()

		var value []uint8

		for id := 0; id < b.N; id++ {
			if n, err := f.ReadAt(value, 0); err != nil {
				log.Fatal(err)
			} else if n < len(value) {
				log.Fatal(io.ErrShortWrite)
			}
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkWriteOpen(b *testing.B) {

	f, err := os.OpenFile("id", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()

	value := []byte("a")

	for id := 0; id < b.N; id++ {
		if n, err := f.WriteAt(value, 0); err != nil {
			log.Fatal(err)
		} else if n < len(value) {
			log.Fatal(io.ErrShortWrite)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}