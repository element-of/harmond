package color

import (
	"testing"
)

func TestNewColoredText(t *testing.T) {
	c, err := NewColoredText(nil, "test")

	if err != nil {
		t.Fatal(err)
	}

	if c.Background != None {
		t.Fatalf("Expected c.Background == None, got %s", c.Background.String())
	}

	if c.Foreground != None {
		t.Fatalf("Expected c.Foreground == None, got %s", c.Foreground.String())
	}
}

func TestNewColoredTextWithForegroundColor(t *testing.T) {
	c, err := NewColoredText([]Color{LightRed}, "this text is red")

	if err != nil {
		t.Fatal(err)
	}

	if c.Background != None {
		t.Fatalf("Expected c.Background == None, got %s", c.Background.String())
	}

	if c.Foreground != LightRed {
		t.Fatalf("Expected c.Foreground == LightRed, got %s", c.Foreground.String())
	}
}

func TestNewColoredTextWithBackgroundAndForegroundColor(t *testing.T) {
	c, err := NewColoredText([]Color{LightRed, Blue}, "this text will blind you")

	if err != nil {
		t.Fatal(err)
	}

	if c.Background != Blue {
		t.Fatalf("Expected c.Background == Blue, got %s", c.Background.String())
	}

	if c.Foreground != LightRed {
		t.Fatalf("Expected c.Foreground == LightRed, got %s", c.Foreground.String())
	}
}

func TestNewColoredTextWithColorsThatMakeNoSense(t *testing.T) {
	_, err := NewColoredText([]Color{LightBlue, Black, Green}, "Hi")
	if err != ErrWrongColorNumber || err == nil {
		t.Fatal("Expected an error")
	}

	_, err = NewColoredText([]Color{Color(55)}, "test")
	if err != ErrBadColor || err == nil {
		t.Fatal(err)
	}
}
