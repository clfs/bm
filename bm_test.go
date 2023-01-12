package bm

import "testing"

func BenchmarkIsEmailGlobalRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmailGlobalRegexp("foo.bar+baz@qux.example.com")
	}
}

func BenchmarkIsEmailChainedRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmailChainedRegexp("foo.bar+baz@qux.example.com")
	}
}

func BenchmarkIsEmailVarRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmailVarRegexp("foo.bar+baz@qux.example.com")
	}
}
