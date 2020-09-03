/*
Package slugify generate slug from unicode string, check Format() for details.
Example:
	package main

	import (
		"fmt"

		"github.com/owarai/slugify"
	)

	func main() {
			text := slugify.Format("hello, 你好，world! 世界！", true)
			fmt.Println(text) // Will print: "hello-你好-world-世界"

			someText := slugify.Format("hello, 你好，world! 世界！", false)
			fmt.Println(someText) // Will print: "hello-world"
	}
*/
package slugify

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/unicode/norm"
)

var (
	// \p{L} or \p{Letter}: any kind of letter from any language.
	// \p{Z} or \p{Separator}: any kind of whitespace or invisible separator.
	// \p{P} or \p{Punctuation}: any kind of punctuation character.

	asciiDecoder           *encoding.Decoder
	asciiEncoder           *encoding.Encoder
	retainGeneralCharacter = regexp.MustCompile(`[^\p{L}\p{Z}\p{P}]`)
	toDash                 = regexp.MustCompile(`[\p{Z}\p{P}]+`)
)

func forceASCII(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if r <= unicode.MaxASCII {
			builder.WriteByte(byte(r))
		}
	}
	res, _ := asciiEncoder.String(builder.String())
	res, _ = asciiDecoder.String(res)
	return res
}

// Format slugify the input(must be valid utf8 string), convert to ASCII if 'allowUnicode' is false(non-ascii character will be removed in this situation).
// Convert spaces or repeated dashes to single dashes. Remove characters that aren't letters, punctuations or separators.
// Convert to lowercase and replace other(may repeated) punctuations or separators to single dash.
func Format(s string, allowUnicode bool) string {
	var value string
	if allowUnicode {
		value = norm.NFKC.String(s)
	} else {
		value = forceASCII(norm.NFKD.String(s))
	}
	value = retainGeneralCharacter.ReplaceAllString(strings.ToLower(value), "")
	value = strings.Trim(toDash.ReplaceAllString(value, "-"), "-")
	return value
}

func init() {
	provider, err := ianaindex.IANA.Encoding("US-ASCII")
	if err != nil {
		panic("package init failed, can not find ascii encode/decoder")
	}
	asciiEncoder = provider.NewEncoder()
	asciiDecoder = provider.NewDecoder()
}
