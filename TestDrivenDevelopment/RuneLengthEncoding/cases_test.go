/*
Implement run-length encoding and decoding.

Run-length encoding (RLE) is a simple form of data compression, where runs
(consecutive data elements) are replaced by just one data value and count.

For example we can represent the original 53 characters with only 13.

```text
"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"  ->  "12WB12W3B24WB"
```

RLE allows the original data to be perfectly reconstructed from
the compressed data, which makes it a lossless data compression.

```text
"AABCCCDEEEE"  ->  "2AB3CD4E"  ->  "AABCCCDEEEE"
```

For simplicity, you can assume that the unencoded string will only contain
the letters A through Z (either lower or upper case) and whitespace. This way
data to be encoded will never contain any numbers and numbers inside data to
be decoded always represent the count for the following character.
*/

package main

var test_cases = []struct {
	input       string
	expected    string
	description string
}{
	{
		"5533",
		"error: digits not allowed",
		"digits",
	},
	{
		"    ",
		"4 ",
		"all whitespaces",
	},
	{
		"",
		"",
		"empthy string",
	},
	{
		"!!!",
		"error special character",
		"special characters",
	},
	{
		"XYZ",
		"XYZ",
		" All non repeating letters",
	},
	{
		"Aeee@@rRR##",
		"A3E2@3R2#",
		"selected special characters",
	},
	{
		"AABBBCCCC",
		"2A3B4C",
		"string with no single characters",
	},
	{
		"A",
		"A",
		"single character",
	},
	{
		"  heee  heee  ",
		"2 H3E2 H3E2 ",
		"white spaces",
	},
	{
		"AaaaaBBeee",
		"5A2B3E",
		"Mix of upper case and lower case",
	},
	{
		"AAA   eee",
		"3A3 3E",
		"mix of upper case, lower case and white spaces",
	},
	{
		"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB",
		"12WB12W3B24WB",
		" Mixed: repeating and non-repeating letters",
	},
}
