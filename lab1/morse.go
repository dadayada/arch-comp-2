package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

var letters = map[[2]byte]rune{
	[2]byte{0x0, 0x17}:  'a',
	[2]byte{0x1, 0xD5}:  'b',
	[2]byte{0x7, 0x5D}:  'c',
	[2]byte{0x0, 0x75}:  'd',
	[2]byte{0x0, 0x1}:   'e',
	[2]byte{0x1, 0x5D}:  'f',
	[2]byte{0x1, 0xDD}:  'g',
	[2]byte{0x0, 0x55}:  'h',
	[2]byte{0x0, 0x5}:   'i',
	[2]byte{0x17, 0x77}: 'j',
	[2]byte{0x1, 0xD7}:  'k',
	[2]byte{0x1, 0x75}:  'l',
	[2]byte{0x0, 0x77}:  'm',
	[2]byte{0x0, 0x1D}:  'n',
	[2]byte{0x7, 0x77}:  'o',
	[2]byte{0x5, 0xDD}:  'p',
	[2]byte{0x1D, 0xD7}: 'q',
	[2]byte{0x0, 0x5D}:  'r',
	[2]byte{0x0, 0x15}:  's',
	[2]byte{0x0, 0x7}:   't',
	[2]byte{0x0, 0x57}:  'u',
	[2]byte{0x1, 0x57}:  'v',
	[2]byte{0x1, 0x77}:  'w',
	[2]byte{0x7, 0x57}:  'x',
	[2]byte{0x1D, 0x77}: 'y',
	[2]byte{0x7, 0x75}:  'z',
	[2]byte{0xFF, 0xFF}: ' ',
}

func decode(r io.Reader) string {
	var tmpBuf [2]byte
	strBuf := bytes.NewBuffer(make([]byte, 0))
	stream := bufio.NewReader(r)
	_, err := stream.Read(tmpBuf[:])
	for err != io.EOF {
		if err != nil {
			log.Fatal(err)
		}
		strBuf.WriteRune(letters[tmpBuf])
		_, err = stream.Read(tmpBuf[:])
	}
	return strBuf.String()
}

func encode(str string) []byte {
	str = strings.ToLower(str)
	var buffer = bytes.NewBuffer(make([]byte, 0))
	for _, letter := range str {
		for key, val := range letters {
			if letter == val {
				buffer.Write(key[:])
			}
		}
	}
	return buffer.Bytes()
}

func createTestData() []byte {
	var buffer = bytes.NewBuffer(make([]byte, 0))
	for i := 0; i < 1000000; i++ {
		buffer.Write(encode("apple "))
	}
	return buffer.Bytes()
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	/*
		f, _ := os.OpenFile("data.dat", os.O_APPEND|os.O_WRONLY, 0777)
		for i := 0; i < 100; i++ {
			_, err := f.Write(createTestData())
			if err != nil {
				log.Fatal(err)
			}
		}
		f.Close()
	*/
	data, _ := ioutil.ReadFile("data.dat")
	decode(bytes.NewReader(data))
}
