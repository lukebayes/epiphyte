package ui_test

import (
	"github.com/waybeams/assert"
	"clock"
	"testing"
	"ui"
	"uiold/context"
)

func TestContext(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		c := context.New()
		assert.NotNil(c)
	})

	t.Run("Sets defaults", func(t *testing.T) {
		c := context.New()
		assert.NotNil(c.Clock())
		assert.NotNil(c.Builder())
	})

	t.Run("Accepts Builder", func(t *testing.T) {
		b := ui.NewBuilder()
		c := context.New(context.Builder(b))
		assert.Equal(b, c.Builder())
	})

	t.Run("Accepts Clock", func(t *testing.T) {
		ck := clock.NewFake()
		c := context.New(context.Clock(ck))
		assert.Equal(ck, c.Clock())
	})

	t.Run("Accepts Clock and Builder", func(t *testing.T) {
		ck := clock.NewFake()
		b := ui.NewBuilder()
		c := context.New(context.Builder(b), context.Clock(ck))
		assert.Equal(c.Clock(), ck)
		assert.Equal(c.Builder(), b)
	})

	t.Run("Adds Font", func(t *testing.T) {
		c := context.New(context.Font("Roboto", "../../testdata/Roboto-Regular.ttf"))
		f := c.Font("Roboto")
		assert.NotNil(f)
		f.SetSize(36)
		w, _ := f.Bounds("ABCD")
		assert.Equal(w, 79)
	})
}
