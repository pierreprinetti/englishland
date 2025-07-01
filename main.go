package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// transcribe converts an Icelandic word into an English-readable phonetic approximation.
// It works by iterating through the word and applying a series of pronunciation rules.
func transcribe(word string) string {
	// Pre-process the word: trim whitespace and convert to lowercase.
	word = strings.ToLower(strings.TrimSpace(word))
	if word == "" {
		return ""
	}

	// Using a strings.Builder is more efficient for building strings in a loop.
	var result strings.Builder
	// Convert the string to a slice of runes to handle multi-byte characters correctly.
	runes := []rune(word)
	n := len(runes)

	for i := 0; i < n; {
		// Add a hyphen between syllables for readability, but not at the beginning.
		if result.Len() > 0 {
			result.WriteString("-")
		}

		// --- Rule Set ---
		// The logic checks for longer phonetic combinations first (2 characters)
		// before falling back to single-character rules.

		// Check for two-character combinations (digraphs)
		if i+1 < n {
			digraph := string(runes[i : i+2])
			switch digraph {
			case "au":
				result.WriteString("oy")
				i += 2
				continue
			case "ey", "ei":
				result.WriteString("ay")
				i += 2
				continue
			case "hv": // Pronounced 'kv'
				result.WriteString("kv")
				i += 2
				continue
			case "ll": // A famous Icelandic sound, approximated as 'tl'
				result.WriteString("tl")
				i += 2
				continue
			case "fn":
				result.WriteString("pn")
				i += 2
				continue
			case "rl":
				result.WriteString("rtl")
				i += 2
				continue
			case "rn":
				result.WriteString("rtn")
				i += 2
				continue
			}
		}

		// If no digraph rule matched, check for single characters (monographs)
		char := runes[i]
		switch char {
		// Vowels with accents
		case 'á':
			result.WriteString("ow")
		case 'é':
			result.WriteString("yeh")
		case 'í', 'ý':
			result.WriteString("ee")
		case 'ó':
			result.WriteString("oh")
		case 'ú':
			result.WriteString("oo")
		case 'æ':
			result.WriteString("eye")
		case 'ö':
			result.WriteString("ur")
		// Special consonants
		case 'þ': // "Thorn" - unvoiced 'th' like in "thin"
			result.WriteString("th")
		case 'ð': // "Eth" - voiced 'th' like in "the"
			result.WriteString("th")
		case 'j':
			result.WriteString("y")
		// Consonants with special rules
		case 'g':
			// 'g' is often silent or soft between vowels or at the end of a word.
			// This is a simplification; the rules are complex. Here we make it a hard 'g'.
			result.WriteString("g")
		case 'f':
			// 'f' sounds like 'v' between vowels
			if i > 0 && i < n-1 && isVowel(runes[i-1]) && isVowel(runes[i+1]) {
				result.WriteString("v")
			} else {
				result.WriteString("f")
			}
		// Standard vowels and consonants
		case 'a':
			result.WriteString("a")
		case 'e':
			result.WriteString("eh")
		case 'i', 'y':
			result.WriteString("i")
		case 'o':
			result.WriteString("o")
		case 'u':
			result.WriteString("uh")
		default:
			// If the character doesn't have a specific rule, append it as is.
			result.WriteRune(char)
		}
		i++
	}

	return result.String()
}

// isVowel is a helper function to check if a rune is a vowel.
// It includes all Icelandic vowel forms.
func isVowel(r rune) bool {
	return strings.ContainsRune("aáeéiíóúüyýæö", r)
}

// main is the entry point of the program.
// It prompts the user for an Icelandic word and prints its phonetic approximation.
func main() {
	fmt.Println("Icelandic Pronunciation Approximator")
	fmt.Println("Enter an Icelandic word (or 'exit' to quit):")

	// Read input from the user in a loop.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			fmt.Println("Bless bless!") // "Bye bye!" in Icelandic
			break
		}

		if input != "" {
			pronunciation := transcribe(input)
			fmt.Printf("English approximation: %s\n\n", pronunciation)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

