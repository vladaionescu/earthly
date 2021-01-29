package conslogging

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/fatih/color"
)

// ColorMode is the mode in which colors are represented in the output.
type ColorMode int

const (
	// AutoColor automatically detects the presence of a TTY to decide if
	// color should be used.
	AutoColor ColorMode = iota
	// NoColor disables use of color.
	NoColor
	// ForceColor forces use of color.
	ForceColor
)

const (
	// NoPadding means the old behavior of printing the full target only.
	NoPadding int = -1
	// DefaultPadding always prints 20 characters for the target, right
	// justified. If it is longer, it prints the right 20 characters.
	DefaultPadding int = 20
)

var currentConsoleMutex sync.Mutex

// ConsoleLogger is a writer for consoles.
type ConsoleLogger struct {
	prefix string
	// metadataMode are printed in a different color.
	metadataMode bool
	// salt is a salt used for color consistency
	// (the same salt will get the same color).
	salt      string
	colorMode ColorMode
	isCached  bool
	isFailed  bool

	// The following are shared between instances and are protected by the mutex.
	mu             *sync.Mutex
	saltColors     map[string]*color.Color
	nextColorIndex *int
	outW           io.Writer
	errW           io.Writer
	trailingLine   bool
	prefixPadding  int
}

// Current returns the current console.
func Current(colorMode ColorMode, prefixPadding int) ConsoleLogger {
	return ConsoleLogger{
		outW:           os.Stdout,
		errW:           os.Stderr,
		colorMode:      colorMode,
		saltColors:     make(map[string]*color.Color),
		nextColorIndex: new(int),
		prefixPadding:  prefixPadding,
		mu:             &currentConsoleMutex,
	}
}

func (cl ConsoleLogger) clone() ConsoleLogger {
	return ConsoleLogger{
		outW:           cl.outW,
		errW:           cl.errW,
		prefix:         cl.prefix,
		metadataMode:   cl.metadataMode,
		salt:           cl.salt,
		isCached:       cl.isCached,
		isFailed:       cl.isFailed,
		saltColors:     cl.saltColors,
		colorMode:      cl.colorMode,
		nextColorIndex: cl.nextColorIndex,
		prefixPadding:  cl.prefixPadding,
		mu:             cl.mu,
	}
}

// WithPrefix returns a ConsoleLogger with a prefix added.
func (cl ConsoleLogger) WithPrefix(prefix string) ConsoleLogger {
	ret := cl.clone()
	ret.prefix = prefix
	ret.salt = prefix
	return ret
}

// WithMetadataMode returns a ConsoleLogger with metadata printing mode set.
func (cl ConsoleLogger) WithMetadataMode(metadataMode bool) ConsoleLogger {
	ret := cl.clone()
	ret.metadataMode = metadataMode
	return ret
}

// WithPrefixAndSalt returns a ConsoleLogger with a prefix and a seed added.
func (cl ConsoleLogger) WithPrefixAndSalt(prefix string, salt string) ConsoleLogger {
	ret := cl.clone()
	ret.prefix = prefix
	ret.salt = salt
	return ret
}

// Prefix returns the console's prefix.
func (cl ConsoleLogger) Prefix() string {
	return cl.prefix
}

// WithCached returns a ConsoleLogger with isCached flag set accordingly.
func (cl ConsoleLogger) WithCached(isCached bool) ConsoleLogger {
	ret := cl.clone()
	ret.isCached = isCached
	return ret
}

// WithFailed returns a ConsoleLogger with isFailed flag set accordingly.
func (cl ConsoleLogger) WithFailed(isFailed bool) ConsoleLogger {
	ret := cl.clone()
	ret.isFailed = isFailed
	return ret
}

// PrintSuccess prints the success message.
func (cl ConsoleLogger) PrintSuccess(msg string) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.printBar(successColor, " SUCCESS ", msg)
}

// PrintFailure prints the failure message.
func (cl ConsoleLogger) PrintFailure(msg string) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	cl.printBar(warnColor, " FAILURE ", msg)
}

func (cl ConsoleLogger) printBar(color *color.Color, center, msg string) {
	if msg != "" {
		center = fmt.Sprintf("%s[%s] ", center, msg)
	}

	totalWidth := 80
	eqBar := strings.Repeat("=", (totalWidth-len(center))/2)
	leftBar := eqBar
	rightBar := eqBar

	// Ensure the width is always totalWidth
	if len(center)%2 == 1 {
		rightBar += "="
	}

	cl.color(color).Fprintf(cl.outW, "%s%s%s\n", leftBar, center, rightBar)
}

// Warnf prints a warning message in red to errWriter
func (cl ConsoleLogger) Warnf(format string, args ...interface{}) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	c := cl.color(warnColor)
	text := fmt.Sprintf(format, args...)
	text = strings.TrimSuffix(text, "\n")

	for _, line := range strings.Split(text, "\n") {
		cl.printPrefix(true)
		c.Fprintf(cl.errW, "%s\n", line)
	}
}

// Printf prints formatted text to the console.
func (cl ConsoleLogger) Printf(format string, args ...interface{}) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	c := cl.color(noColor)
	if cl.metadataMode {
		c = cl.color(metadataModeColor)
	}
	text := fmt.Sprintf(format, args...)
	text = strings.TrimSuffix(text, "\n")
	for _, line := range strings.Split(text, "\n") {
		cl.printPrefix(false)
		c.Fprintf(cl.outW, "%s", line)
		// Don't use a background color for \n.
		cl.color(noColor).Fprintf(cl.outW, "\n")
	}
}

// PrintBytes prints bytes directly to the console.
func (cl ConsoleLogger) PrintBytes(data []byte) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	c := cl.color(noColor)
	if cl.metadataMode {
		c = cl.color(metadataModeColor)
	}

	output := make([]byte, 0, len(data))
	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		ch := data[:size]
		data = data[size:]
		switch r {
		case '\r':
			output = append(output, ch...)
			cl.trailingLine = false
		case '\n':
			output = append(output, ch...)
			cl.trailingLine = false
		default:
			if !cl.trailingLine {
				if len(output) > 0 {
					c.Fprintf(cl.outW, "%s", string(output))
					output = output[:0]
				}
				cl.printPrefix(false)
				cl.trailingLine = true
			}
			output = append(output, ch...)
		}
	}
	if len(output) > 0 {
		c.Fprintf(cl.outW, "%s", string(output))
		output = output[:0]
	}
}

func (cl ConsoleLogger) printPrefix(useErrWriter bool) {
	var w io.Writer
	if useErrWriter {
		w = cl.errW
	} else {
		w = cl.outW
	}

	// Assumes mu locked.
	if cl.prefix == "" {
		return
	}
	c, found := cl.saltColors[cl.salt]
	if !found {
		c = availablePrefixColors[*cl.nextColorIndex]
		cl.saltColors[cl.salt] = c
		*cl.nextColorIndex = (*cl.nextColorIndex + 1) % len(availablePrefixColors)
	}
	c = cl.color(c)
	c.Fprintf(w, cl.prettyPrefix())
	if cl.isFailed {
		w.Write([]byte(" *"))
		cl.color(warnColor).Fprintf(w, "failed")
		w.Write([]byte("*"))
	}
	w.Write([]byte(" | "))
	if cl.isCached {
		w.Write([]byte("*"))
		cl.color(cachedColor).Fprintf(w, "cached")
		w.Write([]byte("* "))
	}
}

func (cl ConsoleLogger) color(c *color.Color) *color.Color {
	switch cl.colorMode {
	case NoColor:
		return noColor
	case ForceColor:
		return c
	case AutoColor:
		if color.NoColor {
			return noColor
		}
		return c
	}
	return noColor
}

var bracketsRegexp = regexp.MustCompile("\\(([^\\]]*)\\)")

func (cl ConsoleLogger) prettyPrefix() string {
	if cl.prefixPadding == NoPadding {
		return cl.prefix
	}

	var brackets string
	bracketParts := strings.SplitN(cl.prefix, "(", 2)
	if len(bracketParts) > 1 {
		brackets = fmt.Sprintf("(%s", bracketParts[1])
	}
	prettyPrefix := bracketParts[0]
	if len(cl.prefix) > cl.prefixPadding {
		parts := strings.Split(cl.prefix, "/")
		target := parts[len(parts)-1]

		truncated := ""
		for _, part := range parts[:len(parts)-1] {
			letter := part
			if len(part) > 0 && part != ".." {
				letter = string(part[0])
			}

			truncated += letter + "/"
		}

		prettyPrefix = truncated + target
	}

	formatString := fmt.Sprintf("%%%vv", cl.prefixPadding)
	return fmt.Sprintf(formatString, fmt.Sprintf("%s%s", prettyPrefix, brackets))
}
