# Sentence Corrector

Sentence Corrector reads a text file (`sample.txt`), applies various text transformations, and outputs the corrected text into a new file (`result.txt`). Transformations include changing case, converting hexadecimal and binary numbers, and adjusting punctuation.

## Features

- **Case Conversion**: Convert words to uppercase, lowercase, or capitalize them.
- **Hexadecimal and Binary Conversion**: Convert hexadecimal and binary numbers to decimal.
- **Punctuation Adjustment**: Correct the spacing around punctuation marks.
- **Article Adjustment**: Change "a" to "an" when followed by a vowel or 'h'.

### Input File Format

The input file (`sample.txt`) should contain sentences that need correction. For example:

```bash
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
```

### Output File Format

The output file (`result.txt`) will contain the corrected sentences. For example:

```bash
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
```

### Running the Program

To run the program, execute the `main.go` file:

```bash
go run main.go sample.txt result.txt
```

Ensure that `sample.txt` is in the same directory as `main.go`. The corrected text will be written to `result.txt`.

## Functions

`hex(str string) string`: Converts a hexadecimal string to decimal.

`bin(str string) string`: Converts a binary string to decimal.

`up(str string) string`: Converts a string to uppercase.

`low(str string) string`: Converts a string to lowercase.

`cap(str string) string`: Capitalizes each word in a string.

`Punctuations(value string) string`: Adjusts spacing around punctuation marks.

`ChangeA(s []string) []string`: Changes "a" to "an" when followed by a vowel or 'h'.

`FixAgain(res string) string`: Fixes additional spacing issues after punctuation.
