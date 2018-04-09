package display

import (
	"assert"
	"testing"
)

type FakeComponent struct {
	Component
}

func NewFake() Displayable {
	return &FakeComponent{}
}

// Create a new factory using our component creation function reference.
var Fake = NewComponentFactory("Fake", NewFake)

func TestComponentFactory(t *testing.T) {
	t.Run("Default State", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		// These two assertions don't appear to be passing my custom equality check. :barf:
		if box.HAlign() != AlignLeft {
			t.Error("Expected AlignLeft, but got: %v", box.HAlign())
		}
		// These two assertions don't appear to be passing my custom equality check. :barf:
		if box.LayoutType() != StackLayoutType {
			t.Error("Expected StackLayout")
		}
		// Width and Height are inferred to zero on request. Clients can ask for StaticWidth and Height
		// for the explicitly configured value.
		assert.Equal(t, box.Height(), 0.0, "GetHeight is derived to zero")
		assert.Equal(t, box.Width(), 0.0, "GetWidth is derived to zero")

		assert.Equal(t, box.ActualHeight(), -1.0, "ActualHeight")
		assert.Equal(t, box.ActualWidth(), -1.0, "ActualWidth")
		assert.Equal(t, box.FlexHeight(), -1.0, "GetFlexHeight")
		assert.Equal(t, box.FlexWidth(), -1.0, "GetFlexWidth")
		assert.Equal(t, box.MaxHeight(), -1.0, "GetMaxHeight")
		assert.Equal(t, box.MaxWidth(), -1.0, "GetMaxWidth")
		assert.Equal(t, box.MinHeight(), -1.0, "GetMinHeight")
		assert.Equal(t, box.MinWidth(), -1.0, "GetMinWidth")
		assert.Equal(t, box.Padding(), -1.0, "GetPadding")
		assert.Equal(t, box.PaddingBottom(), -1.0)
		/*
			assert.Equal(t, box.GetPaddingLeft(), -1.0)
			assert.Equal(t, box.GetPaddingRight(), -1.0)
			assert.Equal(t, box.GetPaddingTop(), -1.0)
			assert.Equal(t, box.GetPrefHeight(), -1.0)
			assert.Equal(t, box.GetPrefWidth(), -1.0)
			assert.Equal(t, box.GetVAlign(), AlignTop)
			assert.Equal(t, box.GetX(), -1.0)
			assert.Equal(t, box.GetY(), -1.0)
			assert.Equal(t, box.GetZ(), -1.0)
			assert.Equal(t, box.GetWidth(), -1.0)
		*/
	})

	t.Run("No Builder", func(t *testing.T) {
		box, _ := Box(NewBuilder(), ID("root"), Children(func(b Builder) {
			Box(b, ID("one"))
			Box(b, ID("two"))
		}))
		if box.ID() != "root" {
			t.Error("Expected a configured Box component")
		}
	})

	t.Run("Child with no builder should fail", func(t *testing.T) {
		unexpectedReslt, err := Box(nil)

		if unexpectedReslt != nil {
			t.Error("Should not have returned a component with no Builder")
		}
		if err == nil {
			t.Error("Expected an error when no component was provided")
		}
	})

	t.Run("Custom type", func(t *testing.T) {
		fake, _ := Fake(NewBuilder())
		if fake == nil {
			t.Error("Expected builder to return new component")
		}
	})

	t.Run("Padding", func(t *testing.T) {
		sprite, _ := Box(NewBuilder(), Padding(10))

		if sprite.Padding() != 10 {
			t.Error("Expected option to set padding")
		}
		if sprite.HorizontalPadding() != 20 {
			t.Error("Expected Padding to update HorizontalPadding")
		}
		if sprite.VerticalPadding() != 20 {
			t.Error("Expected Padding to update VerticalPadding")
		}
		if sprite.PaddingBottom() != 10 {
			t.Error("Expected Padding to update PaddingBottom")
		}
		if sprite.PaddingLeft() != 10 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if sprite.PaddingRight() != 10 {
			t.Error("Expected Padding to update PaddingRight")
		}
		if sprite.PaddingTop() != 10 {
			t.Error("Expected Padding to update PaddingTop")
		}
	})

	t.Run("Padding with specifics", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Padding(10), PaddingLeft(15))
		if box.VerticalPadding() != 20 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.HorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.PaddingLeft() != 15 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if box.PaddingRight() != 10 {
			t.Error("Expected Padding to update PaddingRight")
		}
	})

	t.Run("Padding with specifics is NOT order dependent", func(t *testing.T) {
		box, _ := Box(NewBuilder(), PaddingLeft(15), Padding(10))

		if box.HorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
	})

	t.Run("Padding with specifics will NOT clobber a ZERO setting", func(t *testing.T) {
		box, _ := Box(NewBuilder(), PaddingLeft(0), Padding(10))

		if box.PaddingLeft() != 0 {
			t.Error("Padding option should not clobber a previously set value of Zero")
		}

		if box.HorizontalPadding() != 10 {
			t.Error("Expected zero value padding left to be respected")
		}

		if box.VerticalPadding() != 20 {
			t.Error("Padding should apply to both axis")
		}
	})

	t.Run("Specific Paddings", func(t *testing.T) {
		box, _ := Box(NewBuilder(), PaddingBottom(1), PaddingRight(2), PaddingLeft(3), PaddingTop(4))

		if box.VerticalPadding() != 5 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.HorizontalPadding() != 5 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.PaddingLeft() != 3 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if box.PaddingRight() != 2 {
			t.Error("Expected Padding to update PaddingRight")
		}
		if box.PaddingTop() != 4 {
			t.Error("Expected Padding to update PaddingTop")
		}
		if box.PaddingBottom() != 1 {
			t.Error("Expected Padding to update PaddingBottom")
		}
	})

	t.Run("NewComponentFactoryFrom", func(t *testing.T) {
		Fake := NewComponentFactoryFrom("Fake", VBox, Width(100))

		instance, _ := Fake(NewBuilder(), ID("root"))
		assert.Equal(t, instance.ID(), "root")
		assert.Equal(t, instance.Width(), 100)
	})
}
