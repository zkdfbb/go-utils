package utils

import (
	"fmt"
	"testing"
)

func TestRandomIP(t *testing.T) {
	t.Log(RandomIP())
}

func TestCPUID(t *testing.T) {
	t.Logf("cpuid: %s", CPUID())
}

func TestGetLocalIP(t *testing.T) {
	fmt.Println(GetLocalIP())
}
