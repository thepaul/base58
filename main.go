// base58 utility by paul cannon <p@thepaul.org>

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/btcsuite/btcutil/base58"
	"github.com/spf13/pflag"
)

var (
	decoding = pflag.BoolP("-decode", "d", false,
		"Decode instead of encode")
	check = pflag.BoolP("-check", "c", false,
		"Add a checksum (or, if decoding, expect and check a checksum)")
	version = pflag.Int8P("-version", "v", 0,
		"Use the given version byte when encoding with checksum")
)

func decodeAll(r io.Reader, w io.Writer, check bool) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		var result []byte
		if check {
			var err error
			result, _, err = base58.CheckDecode(word)
			if err != nil {
				return err
			}
		} else {
			result = base58.Decode(word)
		}
		_, err := w.Write(result)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func encodeAll(r io.Reader, w io.Writer, check bool, version int8) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	var result []byte
	if check {
		result = []byte(base58.CheckEncode(data, byte(version)))
	} else {
		result = []byte(base58.Encode(data))
	}
	_, err = w.Write(result)
	return err
}

func main() {
	pflag.Parse()

	var err error
	if *decoding {
		err = decodeAll(os.Stdin, os.Stdout, *check)
	} else {
		err = encodeAll(os.Stdin, os.Stdout, *check, *version)
	}
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		if err != nil {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
