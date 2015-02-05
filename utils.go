package franco

import (
	"golang.org/x/exp/utf8string"
	"regexp"
	"strings"
	"unicode/utf8"
)

func clean(value string) string {
	re := regexp.MustCompile("[\u0021-\u0040]+")
	value = re.ReplaceAllString(value, " ")

	re = regexp.MustCompile(`\s+`)
	value = re.ReplaceAllString(value, " ")

	value = strings.ToLower(value)
	value = strings.Trim(value, " ")

	return value
}

func getTrigrams(value string) []string {
	value = " " + clean(value) + " "
	res := make([]string, 0)
	val := utf8string.NewString(value)
	i := 0
	for i+3 < utf8.RuneCountInString(value) {
		res = append(res, val.Slice(i, i+3))
		i++
	}

	return res
}

func getTrigramsAsMap(value string) map[string]int {
	trigrams := getTrigrams(value)
	res := make(map[string]int)
	for _, t := range trigrams {
		if _, ok := res[t]; !ok {
			res[t] = 0
		}
		res[t]++
	}

	return res
}

func getTrigramsAsTuples(value string) []tuple {
	trigrams := getTrigramsAsMap(value)
	res := make([]tuple, len(trigrams))
	i := 0
	for code, count := range trigrams {
		res[i] = tuple{Code: code, Count: float64(count)}
		i++
	}

	return res
}
