# base58

Base58 encoding and decoding, optionally with checksums

## Usage

By default, encodes all input into base58 with no checksum and outputs the result.

With `-d`/`--decode`, decodes all input _from_ base58 and outputs the result.

With `-x`/`--hex`, expects hexadecimal input (or, if decoding, produces hexadecimal output).

With `-c`/`--check`, add a checksum (or, if decoding, expect and check a checksum).

With `-v N`/`--version N`, use the given version number when encoding with checksum (default 0).

## Examples

```
$ echo hello | base58
tzCkV5DK

$ echo af2c42003efc826ab4361f73f9d890942146fe0ebe806786f8e7190800000000 | base58 -x
CnoVMp6EPFgVAM1QGK2riym9T9GGWybwEZkXcvEAz5qy

$ echo af2c42003efc826ab4361f73f9d890942146fe0ebe806786f8e7190800000000 | base58 -cx  # with checksum
12L9ZFwhzVpuEKMUNUqkaTLGzwY9G24tbiigLiXpmZWKwmcNDDs

$ base58 -d <<<12L9ZFwhzVpuEKMUNUqkaTLGzwY9G24tbiigLiXpmZWKwmcNDDs | od -tx1
0000000    00  af  2c  42  00  3e  fc  82  6a  b4  36  1f  73  f9  d8  90
0000020    94  21  46  fe  0e  be  80  67  86  f8  e7  19  08  00  00  00
0000040    00  6a  35  da  12
0000045

$ base58 -dx <<<12L9ZFwhzVpuEKMUNUqkaTLGzwY9G24tbiigLiXpmZWKwmcNDDs
00af2c42003efc826ab4361f73f9d890942146fe0ebe806786f8e71908000000006a35da12

$ base58 -dcx <<<12L9ZFwhzVpuEKMUNUqkaTLGzwY9G24tbiigLiXpmZWKwmcNDDs
af2c42003efc826ab4361f73f9d890942146fe0ebe806786f8e7190800000000

$ echo tzCkV5DK | base58 -d
hello
```
