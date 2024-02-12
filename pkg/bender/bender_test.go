package bender_test

import (
	"mp/benderhash/pkg/bender"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	testCases := []struct {
		key           string
		value         string
		wantInsertErr error
		wantGetErr    error
	}{
		{
			key:   "grape tannins and C₂H₅OH",
			value: "red wine",
		},
		{
			key:           "grape tannins and sugar",
			value:         "you know, for kids",
			wantInsertErr: bender.ErrPleaseInsertLiquor,
			wantGetErr:    bender.ErrNotFound,
		},
		{
			key:   "Motor oil, EtOH, and carbon monoxide",
			value: "Olde Fortran Malt Liquor",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.key, func(t *testing.T) {
			d := bender.Hash{}

			err := d.Insert(tC.key, tC.value)
			assert.Equal(t, tC.wantInsertErr, err)

			got, err := d.Get(tC.key)
			assert.Equal(t, tC.wantGetErr, err)
			if err == nil {
				assert.Equal(t, tC.value, got)
			}
		})
	}
}
