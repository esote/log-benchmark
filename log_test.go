package benchmark

import (
	"bytes"
	"log"
	"os"
	"testing"
)

const (
	msg = "string message"
)

func BenchmarkBuffer(b *testing.B) {
	var buf bytes.Buffer
	l := log.New(&buf, "", log.LstdFlags)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = l.Output(2, msg)
	}
}

func BenchmarkFile(b *testing.B) {
	const filename = "file-test"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0200)

	if err != nil {
		log.Fatal(err)
	}

	if err = f.Truncate(0); err != nil {
		log.Fatal(err)
	}

	if _, err = f.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	l := log.New(f, "", log.LstdFlags)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = l.Output(2, msg)
	}

	b.StopTimer()

	f.Truncate(0)

	f.Close()
}

func BenchmarkNull(b *testing.B) {
	f, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0200)

	if err != nil {
		log.Fatal(err)
	}

	l := log.New(f, "", log.LstdFlags)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = l.Output(2, msg)
	}

	b.StopTimer()

	f.Close()
}
