package main

import (
	"bufio"
	"io"
)

func fromHex(r io.Reader) *decodeHexReader {
	return &decodeHexReader{
		source: bufio.NewReader(r),
	}
}

type decodeHexReader struct {
	source *bufio.Reader
}

func (r *decodeHexReader) Read(buf []byte) (n int, err error) {
	for i := 0; i < len(buf); i++ {
		var nybbles [2]int8
		var ok bool
		for nybbleNum := 0; nybbleNum < 2; {
			b, err := r.source.ReadByte()
			if err != nil {
				return i, err
			}
			nybbles[nybbleNum], ok = hexValue(b)
			if !ok {
				if isSpace(b) && nybbleNum == 0 {
					// skip whitespace, but only between bytes
					continue
				}
				return i, InvalidInputErr
			}
			nybbleNum++
		}
		buf[i] = byte(nybbles[0]*16 + nybbles[1])
	}
	return len(buf), nil
}

func hexValue(b byte) (n int8, ok bool) {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return int8(b - '0'), true
	case 'A', 'B', 'C', 'D', 'E', 'F':
		return int8(b - 'A' + 10), true
	case 'a', 'b', 'c', 'd', 'e', 'f':
		return int8(b - 'a' + 10), true
	default:
		return 0, false
	}
}

func toHex(w io.Writer) *encodeHexWriter {
	return &encodeHexWriter{
		dest: bufio.NewWriter(w),
	}
}

type encodeHexWriter struct {
	dest *bufio.Writer
}

func (w *encodeHexWriter) Write(buf []byte) (n int, err error) {
	var nybbles [2]byte
	for i, b := range buf {
		nybbles[0] = (b & 0xf0) >> 4
		nybbles[1] = b & 0x0f
		for _, nybble := range nybbles {
			err = w.dest.WriteByte(hexDigit(int8(nybble)))
			if err != nil {
				return i, err
			}
		}
	}
	err = w.dest.Flush()
	return len(buf), err
}

func hexDigit(val int8) byte {
	if val >= 0 && val <= 9 {
		return byte('0' + val)
	}
	if val >= 10 && val <= 15 {
		return byte('a' - 10 + val)
	}
	return byte('X')
}

func isSpace(b byte) bool {
	switch b {
	case ' ', '\t', '\r', '\n', '\v', '\f':
		return true
	}
	return false
}
