// base58 utility by paul cannon <p@thepaul.org>

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
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
	useHex = pflag.BoolP("-hex", "x", false,
		"Expect hexadecimal input (or, if decoding, produce hexadecimal output)")
)

func decodeAllFromBase58(r io.Reader, w io.Writer, check, useHex bool) (err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		var result []byte
		if check {
			result, _, err = base58.CheckDecode(word)
			if err != nil {
				return err
			}
		} else {
			result = base58.Decode(word)
		}
		if useHex {
			_, err = fmt.Fprintf(w, "%x\n", result)
		} else {
			_, err = w.Write(result)
		}
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func encodeAllToBase58(r io.Reader, w io.Writer, check bool, version int8) error {
	data, err := io.ReadAll(r)
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

func encodeAllHexToBase58(r io.Reader, w io.Writer, check bool, version int8) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		data, err := hex.DecodeString(word)
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
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func main() {
	pflag.Parse()

	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	var err error
	if *decoding {
		err = decodeAllFromBase58(input, output, *check, *useHex)
	} else {
		if *useHex {
			err = encodeAllHexToBase58(input, output, *check, *version)
		} else {
			err = encodeAllToBase58(input, output, *check, *version)
		}
	}
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		if err != nil {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
