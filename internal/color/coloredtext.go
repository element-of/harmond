package color

import (
	"errors"
	"fmt"
)

var (
	ErrWrongColorNumber = errors.New("Wrong number of colors")
)

// ColoredText is a section of text that is colored with mIRC color codes.
type ColoredText struct {
	Foreground Color  `json:"foreground"`
	Background Color  `json:"background"`
	Content    string `json:"content"`
}

// IRCString returns a ColoredText as it should be shown on IRC.
func (c ColoredText) IRCString() string {
	if c.Background != None {
		return fmt.Sprintf("%c%d,%d%s%c", ColorCodeDelim, c.Foreground, c.Background, c.Content, ColorCodeDelim)
	} else if c.Foreground != None {
		return fmt.Sprintf("%c%d%s%c", ColorCodeDelim, c.Foreground, c.Content, ColorCodeDelim)
	}

	return c.Content
}

// String returns a ColoredText as a string.
func (c ColoredText) String() string {
	ret := "<ColoredText "

	if c.Foreground != None {
		ret = ret + fmt.Sprintf("Foreground: %s ", c.Foreground.String())
	}

	if c.Background != None {
		ret = ret + fmt.Sprintf("Foreground: %s ", c.Foreground.String())
	}

	ret = ret + fmt.Sprintf(`Content: "%s"`, c.Content)

	return ret
}

// NewColoredText makes a ColoredText with given colors and a content or returns
// a blank ColoredText and an error.
func NewColoredText(colors []Color, content string) (c ColoredText, err error) {
	switch len(colors) {
	case 0:
		c.Background = None
		c.Foreground = None
		c.Content = content

	case 1:
		c.Background = None
		if colors[0].Validate() {
			c.Foreground = colors[0]
		} else {
			return ColoredText{}, ErrBadColor
		}

	case 2:
		if colors[0].Validate() {
			c.Foreground = colors[0]
		} else {
			return ColoredText{}, ErrBadColor
		}

		if colors[1].Validate() {
			c.Background = colors[1]
		} else {
			return ColoredText{}, ErrBadColor
		}

	default:
		return ColoredText{}, ErrWrongColorNumber

	}

	return
}
