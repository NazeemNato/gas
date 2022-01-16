package utils

import "fmt"

func Avatar(hash int, size string) string {
	// creating two gradient colors
	c1 := `<stop stop-color="` + fmt.Sprintf("#%06x", (hash>>24)%0xffffff+(hash>>16)%0xffffff+(hash>>8)%0xffffff) + `" offset="0%"></stop>`
	c2 := `<stop stop-color="` + fmt.Sprintf("#%06x", (hash>>16)%0xffffff+(hash>>8)%0xffffff+(hash>>2)%0xffffff) + `" offset="100%"></stop>`
	// creating the svg
	linearGradient := ` <linearGradient x1="0%" y1="0%" x2="100%" y2="100%" id="g"> ` + c1 + c2 + ` </linearGradient>`
	defs := ` <defs>` + linearGradient + `</defs>`
	g := `<g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd"><rect id="Rectangle" fill="url(#g)" x="0" y="0" width="80" height="80"></rect></g>`
	// returning svg
	return fmt.Sprintf(`<svg width="%vpx" height="%vpx" viewBox="0 0 80 80" version="1.1" xmlns="http://www.w3.org/2000/svg" xlink="http://www.w3.org/1999/xlink">
	%v
	%v
	</svg>`, size, size, defs, g)
}
