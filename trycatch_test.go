package trycatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry_ShouldCatchError(t *testing.T){
	var wasCatched = false

	Try(func() {
		RaiseError(&Error{"Oh no! something went wrong!"})
	}).
	Catch(func(err *Error) {
		wasCatched = true
	}).
	Do()

	assert.True(t, wasCatched)
}

func TestTry_NotCatchedErrorsShouldRaisedUpper(t *testing.T){
	defer func() {
		r := recover()
		assert.True(t, r != nil)
	}()

	Try(func() {
		RaiseError(&Error{"Oh no! something went wrong!"})
	}).
	Finally(func(){

	}).
	Do()
}