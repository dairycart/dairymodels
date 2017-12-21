package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	exampleSKU = "example"
	//exampleTimeString        = "2016-12-01 12:00:00.000000"
	exampleGarbageInput      = `{"things": "stuff"}`
	exampleMarshalTimeString = "2016-12-31T12:00:00.000000Z"
)

func TestDairytimeMarshalText(t *testing.T) {
	t.Parallel()

	out, err := time.Parse("2006-01-02 03:04:00.000000", "2016-12-31 12:00:00.000000")
	require.Nil(t, err)

	expected := []byte(exampleMarshalTimeString)
	example := Dairytime{Time: out}
	actual, err := example.MarshalText()

	require.Nil(t, err)
	require.Equal(t, expected, actual, "Marshaled time string should marshal correctly")
}

func TestDairytimeUnmarshalText(t *testing.T) {
	t.Parallel()
	example := []byte(exampleMarshalTimeString)
	nt := Dairytime{}
	err := nt.UnmarshalText(example)
	require.Nil(t, err)
}
