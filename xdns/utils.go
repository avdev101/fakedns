package xdns

import "regexp"

func getIps(s string) []string {

	var pattern = regexp.MustCompile("[\\d]+\\.[\\d]+\\.[\\d]+\\.[\\d]+")

	return pattern.FindAllString(s, -1)
}
