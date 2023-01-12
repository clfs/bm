// Package bm is a workspace for quick Go benchmarks.
package bm

import "regexp"

var emailRegexp = regexp.MustCompile(`^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$`)

func IsEmailGlobalRegexp(s string) bool {
	return emailRegexp.MatchString(s)
}

func IsEmailChainedRegexp(s string) bool {
	return regexp.MustCompile(`^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$`).MatchString(s)
}

func IsEmailVarRegexp(s string) bool {
	re := regexp.MustCompile(`^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$`)
	return re.MatchString(s)
}
