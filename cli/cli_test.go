package cli

import (
	"sync"
	"testing"
)

func BenchmarkGetUserFromGithub(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		GetUserFromGithub("M-Krishna", &wg)
	}
	wg.Wait()
}
