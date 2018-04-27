package display

import (
	"assert"
	"testing"
)

func TestFocusable(t *testing.T) {
	t.Run("Blurred", func(t *testing.T) {
		instance, _ := Button(NewBuilder(), Blurred())
		assert.False(t, instance.Focused())
	})

	t.Run("Focused", func(t *testing.T) {
		instance, _ := Button(NewBuilder(), Focused())
		assert.True(t, instance.Focused())
	})

	t.Run("Unfocuses previously focused elements", func(t *testing.T) {
		instance, _ := VBox(NewBuilder(), Children(func(b Builder) {
			Button(b, ID("abcd"))
			Button(b, ID("efgh"))
			Button(b, ID("ijkl"))
			Button(b, ID("mnop"))
		}))

		children := instance.Children()
		abcd := children[0].(Focusable)
		efgh := children[1].(Focusable)
		ijkl := children[2].(Focusable)
		mnop := children[3].(Focusable)

		abcd.Focus()
		assert.True(t, abcd.Focused())
		assert.False(t, efgh.Focused())
		assert.False(t, ijkl.Focused())
		assert.False(t, mnop.Focused())

		ijkl.Focus()
		assert.False(t, abcd.Focused())
		assert.False(t, efgh.Focused())
		assert.True(t, ijkl.Focused())
		assert.False(t, mnop.Focused())
	})

	var createTree = func() Displayable {
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			Box(b)
			Box(b)
			Box(b, Children(func() {
				Box(b, ID("uvwx"))
				Button(b, ID("abcd"))
				Box(b, ID("efgh"), IsFocusable(true), Children(func() {
					Box(b, ID("ijkl"), Children(func() {
						Box(b, ID("mnop"))
					}))
				}))
				Button(b, ID("qrst"))
			}))
		}))

		return root
	}

	t.Run("FocusablePath() returns nearest focusable parent", func(t *testing.T) {
		root := createTree()
		child := root.FindComponentByID("mnop")

		nonFocusable := root.FindComponentByID("uvwx")
		assert.Equal(t, nonFocusable.NearestFocusable().Path(), root.Path())

		focusable := root.FindComponentByID("efgh")
		assert.Equal(t, focusable.Path(), focusable.NearestFocusable().Path(), "returns self too")

		expected := child.NearestFocusable()
		assert.Equal(t, focusable.Path(), expected.Path(), "Child returns Focusable grandparent")
	})

	t.Run("Last focusable is blurred", func(t *testing.T) {
		root := createTree()
		abcd := root.FindComponentByID("abcd")
		qrst := root.FindComponentByID("qrst")
		abcd.Focus()
		assert.True(t, abcd.Focused())
		assert.False(t, qrst.Focused())
		assert.Equal(t, root.FocusedChild().Path(), abcd.Path())

		qrst.Focus()
		assert.False(t, abcd.Focused())
		assert.True(t, qrst.Focused())
		assert.Equal(t, root.FocusedChild().Path(), qrst.Path())
	})
}
