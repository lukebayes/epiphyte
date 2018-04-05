package display

import (
	"errors"
)

// BuilderOption is a configuration option for Builders.
type BuilderOption func(b Builder) error

// ComponentComposer is a composition function that components send to the
// Children() option when composing children using an anonymous function.
type ComponentComposer func(b Builder)

// Builder is a basic wrapper around a stack that enables component
// composition.
type Builder interface {
	Push(d Displayable, options ...ComponentOption)
	Peek() Displayable
}

type builder struct {
	root      Displayable
	stack     Stack
	lastError error
	types     map[string]bool
}

func (b *builder) getStack() Stack {
	if b.stack == nil {
		b.stack = NewDisplayStack()
	}
	return b.stack
}

// Current returns the current entry in the Builder stack.
// This method only works while the component declarations are being processed.
func (b *builder) Peek() Displayable {
	return b.getStack().Peek()
}

func (b *builder) callComposeFunctionFor(d Displayable) (err error) {
	composeSimple := d.GetComposeSimple()
	if composeSimple != nil {
		composeSimple()
		return nil
	}
	composeWithBuilder := d.GetComposeWithBuilder()
	if composeWithBuilder != nil {
		composeWithBuilder(b)
		return nil
	}

	return errors.New("no compose function found")
}

// Push accepts a new Displayable to place on the stack and processes the
// optional component composition function if one was provided.
func (b *builder) Push(d Displayable, options ...ComponentOption) {
	if b.root == nil {
		b.root = d
	}

	stack := b.getStack()

	// Get the parent element if one exists
	parent := stack.Peek()

	if parent == nil {
		if b.root != d {
			// It looks like we have a second root definition in the outer factory
			// function
			b.lastError = errors.New("box factory function should only have a single root node")
			return
		}
	} else {
		parent.AddChild(d)
	}

	// Push the element onto the stack
	stack.Push(d)

	for _, option := range options {
		err := option(d)
		if err != nil {
			// If an option error is found, bail with it, for now.
			b.lastError = err
			return
		}
	}

	// Process composition function to build children
	composeError := b.callComposeFunctionFor(d)
	if composeError != nil && b.lastError == nil {
		b.lastError = composeError
	}

	// Pop the element off the stack
	stack.Pop()

	if b.root == d {
		d.Layout()
	}
}

// NewBuilder returns a clean builder instance.
func NewBuilder() Builder {
	return &builder{}
}
