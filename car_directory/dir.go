package carDirectory

import (
	"fmt"
	"os"
	"strings"

	"strconv"

	"github.com/mgutz/ansi"
)

const carPaint = "white:blue"

// Directory car
type Car struct {
	paint string
	// Current directory
	Pwd string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_DIRECTORY_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)

	dir := c.Pwd

	if os.Getenv("HOME") == dir {
		dir = "~"
	} else {
		ps := string(os.PathSeparator)

		max_length := 3
		if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT"); e != "" {
			ml, err := strconv.Atoi(e)
			if err == nil {
				max_length = ml
			}
		}

		dirs := strings.Split(dir, ps)
		if max_length > 0 && len(dirs) > max_length+1 {
			f := len(dirs) - max_length
			p := dirs[f:]
			dir = fmt.Sprintf("...%s", strings.Join(p, ps))
		}
	}

	out <- ansi.Color(fmt.Sprintf("%s", dir), c.GetPaint())
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_DIRECTORY_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_DIRECTORY_SEPARATOR_SYMBOL")
}
