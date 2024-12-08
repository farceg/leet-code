package substringwithconcatenation

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
func FindSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	totalWords := len(words)
	substringLen := wordLen * totalWords
	result := []int{}
	wordCount := make(map[string]int)

	// Contar ocurrencias de cada palabra
	for _, word := range words {
		wordCount[word]++
	}

	for i := 0; i <= len(s)-substringLen; i++ {
		seen := make(map[string]int)
		j := 0

		// Comprobar cada palabra en la ventana
		for j < totalWords {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			if count, ok := wordCount[word]; ok {
				seen[word]++
				if seen[word] > count {
					break
				}
			} else {
				break
			}
			j++
		}

		if j == totalWords {
			result = append(result, i)
		}
	}

	return result
}
