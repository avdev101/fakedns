package xdns

import (
	"regexp"
)

func getIps(s string) []string {

	pattern := regexp.MustCompile("[\\d]+\\.[\\d]+\\.[\\d]+\\.[\\d]+")

	return pattern.FindAllString(s, -1)
}

func getScheme(s string) []int {
	result := make([]int, 0)
	pattern := regexp.MustCompile("s([\\d]+)")
	groups := pattern.FindStringSubmatch(s)

	if len(groups) == 0 {
		return result
	}

	for _, ch := range groups[1] {
		i := int(ch - '0')
		result = append(result, i)
	}

	return result
}
