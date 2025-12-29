package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	want := "Hello Go"

	got := hello()

	require.Equal(t, want, got)
}
