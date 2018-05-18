package controls

import (
	"github.com/waybeams/assert"
	"events"
	"testing"
	"uiold/context"
	. "uiold/opts"
)

func TestButton(t *testing.T) {
	t.Run("Callable", func(t *testing.T) {
		btn := Button(context.New(), Text("Submit"))
		assert.Equal(btn.Text(), "Submit")
	})

	t.Run("Focusable", func(t *testing.T) {
		var calledWith events.Event
		var clickHandler = func(e events.Event) {
			calledWith = e
		}
		button := Box(context.New(), Text("Submit"), OnClick(clickHandler))
		button.Emit(events.New(events.Clicked, button, nil))
		assert.Equal(calledWith.Target(), button)
	})

	t.Run("Uses label dimensions", func(t *testing.T) {
		btn := Button(context.NewTestContext(), TestOptions(), Text("Submit Something"))
		assert.Equal(btn.Width(), 102)
		assert.Equal(btn.Height(), 31)
	})
}
