package main

import (
	"debug"
	"fmt"
	"runtime/debug"

	"github.com/ydm/goinject"
	"github.com/ydm/goinject/dashes-again"
	"github.com/ydm/goinject/nope"
	"github.com/ydm/goinject/package-with-dashes"
	"github.com/ydm/goinject/package-with-dashes/inner"
)

func main() {
	fmt.Printf("goinject.Version=%s\n", goinject.Version)
	fmt.Printf("nope.Version=%s\n", nope.Version)
	fmt.Printf("something_completely_different.Version=%s\n", something_completely_different.Version)
	fmt.Printf("package_with_dashes.Version=%s\n", package_with_dashes.Version)
	fmt.Printf("package_with_dashes.inner.Version=%s\n", blabla.Version)
	fmt.Printf("package_with_dashes.Hello()=%s\n", package_with_dashes.Hello())
	// debug.ReadBuildInfo()
	debug.ReadBuildInfo()
}
