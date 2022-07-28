package strdice

import "strings"

func makeBigrams(s string) map[string]int {
	m := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		bigram := s[i : i+2]
		m[bigram]++
	}

	return m
}

func Compute(s1 string, s2 string) float64 {
	// Replace all spaces
	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")

	// Base cases
	if s1 == s2 {
		return 1
	}
	if len(s1) < 2 || len(s2) < 2 {
		return 0
	}

	// Bigrams for s1 and s2
	bigrams1 := makeBigrams(s1)
	bigrams2 := makeBigrams(s2)

	// Find intersection between two bigram sets
	intersection := 0
	for k, v := range bigrams2 {
		if v < bigrams1[k] {
			intersection += v
		} else {
			intersection += bigrams1[k]
		}
	}

	// No. of bigrams in a string is len(string) - 1
	total := len(s1) + len(s2) - 2

	return (float64(intersection) * 2) / float64(total)
}
