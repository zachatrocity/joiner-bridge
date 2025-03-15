package main

import (
	"maunium.net/go/mautrix/bridgev2/matrix/mxmain"

	"github.com/zachatrocity/joiner-bridge/pkg/connector"
)

// Information to find out exactly which commit the bridge was built from.
// These are filled at build time with the -X linker flag.
var (
	Tag       = "unknown"
	Commit    = "unknown"
	BuildTime = "unknown"
)

func main() {
	m := mxmain.BridgeMain{
		Name:        "joiner-bridge",
		Description: "A Matrix bridge for automatically joining rooms",
		URL:         "https://github.com/zachatrocity/joiner-bridge",
		Version:     "0.1.0",
		Connector:   &connector.RoomJoinerConnector{},
	}
	m.InitVersion(Tag, Commit, BuildTime)
	m.Run()
}
