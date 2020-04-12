package main

import (
	"flag"
	"fmt"

	"github.com/amisini/scraping_portals/portals"
)

var (
	portal = flag.String("portal", "telegrafi", "portal-name")
)

func main() {

	flag.Parse()

	switch *portal {
	case "telegrafi":
		portals.Telegrafi()
	case "gazetaexpress":
		portals.GazetaExpress()
	case "indeksonline":
		portals.IndeksOnline()
	default:
		fmt.Println("Wrong portal")
	}

}
