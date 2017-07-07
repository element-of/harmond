package color

// A color code that mIRC defines
type Color int

//go:generate stringer -type=Color

// Validate makes sure that the recieving color makes sense.
func (c Color) Validate() bool {
	return c >= None && c < (LightGray+1)
}
