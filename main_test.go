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
	from := 380670000000
	b.ResetTimer()
	for id := int64(from); id < int64(1000+from); id++ {
		if _, err := Get(strconv.FormatInt(id, 10), "billing"); err != nil {
			log.Println(err)
		}
	}
}

func BenchmarkWrite(b *testing.B) {
	from := 380670000000
	b.ResetTimer()
	for id := int64(from); id < int64(1000+from); id++ {
		if err := Set(strconv.FormatInt(id, 10), "billing", strconv.FormatInt(rand.Int63n(3), 10)); err != nil {
			log.Println(err)
		}
	}
}

func BenchmarkReadOpen(b *testing.B) {

	if f, err := os.OpenFile("id", os.O_RDONLY|os.O_CREATE, 600); err != nil {
		log.Fatal(err)
	} else {

		var value []uint8

		b.ResetTimer()
		for id := 0; id < b.N; id++ {
			if n, err := f.ReadAt(value, 0); err != nil {
				log.Fatal(err)
			} else if n < len(value) {
				log.Fatal(io.ErrShortWrite)
			}
		}
		b.StopTimer()

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkWriteOpen(b *testing.B) {

	f, err := os.OpenFile("id_", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	value := []byte("a")

	b.ResetTimer()
	for id := 0; id < b.N; id++ {
		if n, err := f.WriteAt(value, 0); err != nil {
			log.Fatal(err)
		} else if n < len(value) {
			log.Fatal(io.ErrShortWrite)
		}
	}
	b.StopTimer()

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	if os.Remove("id_") != nil {
		log.Fatal(err)
	}
}
