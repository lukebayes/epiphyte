package fake

import (
	"math"
)

// Command stores method name and arguments for a given call.
type Command struct {
	Name string
	Args []interface{}
}

// Fake is a drawing surface that is provided to test Draw implementations.
// Rather than rendering into some hardware interface, the methods provided here
// will simply record that they were called and with what arguments.
type Fake struct {
	commands []Command
	width      float64
	height     float64
    pixelRatio float32
}

func (s *Fake) AddFont(name, path string) {
	args := []interface{}{name, path}
	s.commands = append(s.commands, Command{Name: "AddFont", Args: args})
}

func (s *Fake) Init() {
	s.commands = append(s.commands, Command{Name: "Init"})
    s.pixelRatio = 1
}

func (s *Fake) Close() {
	// nooop
}

func (s *Fake) BeginFrame() {
	s.commands = append(s.commands, Command{Name: "BeginFrame"})
}

func (s *Fake) EndFrame() {
	s.commands = append(s.commands, Command{Name: "EndFrame"})
}

// GetCommands returns the collection of commands that have been made against
// a given instance of the Fake.
func (s *Fake) GetCommands() []Command {
	return s.commands
}

// CreateFont creates and caches the font atlas.
func (s *Fake) CreateFont(name, path string) {
	args := []interface{}{name, path}
	s.commands = append(s.commands, Command{Name: "CreateFont", Args: args})
}

func (s *Fake) SetScale(x, y float32) {
}

func (s *Fake) SetPixelRatio(ratio float32) {
    s.pixelRatio = ratio
}

// SetFillColor stores the provided Hex RGBA fill color (e.g., 0xffcc00ff).
func (s *Fake) SetFillColor(color uint) {
	args := []interface{}{color}
	s.commands = append(s.commands, Command{Name: "SetFillColor", Args: args})
}

// SetStrokeColor stores the provided Hex RGBA stroke color (e.g., 0xffcc00ff).
func (s *Fake) SetStrokeColor(color uint) {
	args := []interface{}{color}
	s.commands = append(s.commands, Command{Name: "SetStrokeColor", Args: args})
}

func (s *Fake) MoveTo(x float64, y float64) {
	args := []interface{}{x, y}
	s.commands = append(s.commands, Command{Name: "MoveTo", Args: args})
}

// SetStrokeWidth sets the stroke width
func (s *Fake) SetStrokeWidth(width float64) {
	args := []interface{}{width}
	s.commands = append(s.commands, Command{Name: "SetLineWidth", Args: args})
}

// Stroke draws a stroke around the last created shape.
func (s *Fake) Stroke() {
	s.commands = append(s.commands, Command{Name: "Stroke"})
}

// Arc draws a arc along the provided point, radius and angles.
func (s *Fake) Arc(xc, yc, radius, angle1, angle2 float64) {
	args := []interface{}{xc, yc, radius, angle1, angle2}
	s.commands = append(s.commands, Command{Name: "Arc", Args: args})
}

func (s *Fake) BeginPath() {
	s.commands = append(s.commands, Command{Name: "BeginPath"})
}

func (s *Fake) DebugDumpPathCache() {
	s.commands = append(s.commands, Command{Name: "DebugDumpCachePath"})
}

// Fill will fill the last created shape.
func (s *Fake) Fill() {
	s.commands = append(s.commands, Command{Name: "Fill"})
}

// Rect draws a rectangle on the provided point and width and height.
func (s *Fake) Rect(x, y, width, height float64) {
	args := []interface{}{x, y, width, height}
	s.commands = append(s.commands, Command{Name: "Rect", Args: args})
}

// RoundedRect draws a rectangle with rounded corners on the provided point and width and height.
func (s *Fake) RoundedRect(x, y, width, height, radius float64) {
	args := []interface{}{x, y, width, height, radius}
	s.commands = append(s.commands, Command{Name: "RoundedRect", Args: args})
}

func (s *Fake) SetFontSize(size float64) {
	args := []interface{}{size}
	s.commands = append(s.commands, Command{Name: "SetFontSize", Args: args})
}

func (s *Fake) SetFontFace(face string) {
	args := []interface{}{face}
	s.commands = append(s.commands, Command{Name: "SetFontFace", Args: args})
}

func (s *Fake) Text(x float64, y float64, text string) {
	args := []interface{}{x, y, text}
	s.commands = append(s.commands, Command{Name: "Text", Args: args})
}

func (s *Fake) SetWidth(w float64) {
	s.width = w
	args := []interface{}{w}
	s.commands = append(s.commands, Command{Name: "SetWidth", Args: args})
}

func (s *Fake) SetHeight(h float64) {
	s.height = h
	args := []interface{}{h}
	s.commands = append(s.commands, Command{Name: "SetHeight", Args: args})
}

func (s *Fake) Width() float64 {
	s.commands = append(s.commands, Command{Name: "Width"})
	return s.width
}

func (s *Fake) Height() float64 {
	s.commands = append(s.commands, Command{Name: "Height"})
	return s.height
}

func (s *Fake) TextBounds(face string, size float64, text string) (x, y, w, h float64) {
	args := []interface{}{face, size, text}
	s.commands = append(s.commands, Command{Name: "Text", Args: args})
	// NOTE(lbayes): The following text metrics are fake magic values that obviously
	// will not kern or layout properly, but are sufficient to get boxes in the test
	// environment.
	x = -0.5
	y = -size
	w = math.Floor((size * float64(len(text))) * 0.423)
	h = size
	return x, y, w, h
}

// NewSurface returns a new fake Surface.
func NewSurface() *Fake {
	return &Fake{}
}
