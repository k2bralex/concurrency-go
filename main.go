package main

import "concurrency/md_2"

func main() {

	md_2.LineGenerator()
	md_2.FactorPower(15, 10)
	text := "Golang rune type is the alias for int32, and it is used to indicate than the integer " +
		"represents the code point. ASCII defines 128 characters, identified by the code points 0–127." +
		" When you convert the string to a rune slice, you get the new slice that contains the Unicode " +
		"code points (runes) of a string.\n\nByte slice is just like a string, but mutable. For example, " +
		"you can change each byte or character. This is very efficient for working with file content, " +
		"either as a text file, binary file, or IO stream from networking.\n\nRune slice is like the byte " +
		"slice, except that each index is a character instead of a byte. This is best if you work with " +
		"text files that have lots of non-ASCII characters, such as Chinese text or math formulas ∑ or " +
		"text with emoji ♥."
	md_2.Report(text)

	/*md_1.PrintRange(-10, 50, 5)
	md_1.MinMaxAvgGoroutines(-10000, 10000, 1000000)
	fmt.Println("*================================*")
	md_1.MinMaxAvgNoGoroutines(-10000, 10000, 1000000)*/
}
