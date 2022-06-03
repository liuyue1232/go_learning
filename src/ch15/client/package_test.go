package client

import (
	"ch15/series"
	"ch18"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	t.Log(csp.Hello(5))
}
