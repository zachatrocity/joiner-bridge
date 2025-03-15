package connector

import (
	"context"
	"fmt"

	"go.mau.fi/util/configupgrade"
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/database"
	"maunium.net/go/mautrix/id"
)

// RoomJoinerConnector implements a simple bridge that just joins rooms
type RoomJoinerConnector struct {
	br *bridgev2.Bridge
}

var _ bridgev2.NetworkConnector = (*RoomJoinerConnector)(nil)

func (rc *RoomJoinerConnector) Init(bridge *bridgev2.Bridge) {
	rc.br = bridge
}

func (rc *RoomJoinerConnector) Start(ctx context.Context) error {
	return nil
}

func (rc *RoomJoinerConnector) GetCapabilities() *bridgev2.NetworkGeneralCapabilities {
	return &bridgev2.NetworkGeneralCapabilities{}
}

func (rc *RoomJoinerConnector) GetName() bridgev2.BridgeName {
	return bridgev2.BridgeName{
		DisplayName:      "Room Joiner",
		NetworkURL:       "https://github.com/zachatrocity/joiner-bridge",
		NetworkID:        "joiner",
		BeeperBridgeType: "go.mau.fi/joiner-bridge",
		DefaultPort:      29323,
	}
}

func (rc *RoomJoinerConnector) GetConfig() (example string, data any, upgrader configupgrade.Upgrader) {
	return "", nil, nil
}

func (rc *RoomJoinerConnector) GetDBMetaTypes() database.MetaTypes {
	return database.MetaTypes{}
}

func (rc *RoomJoinerConnector) GetBridgeInfoVersion() (int, int) {
	return 1, 0
}

func (rc *RoomJoinerConnector) GetLoginFlows() []bridgev2.LoginFlow {
	return nil
}

func (rc *RoomJoinerConnector) CreateLogin(ctx context.Context, user *bridgev2.User, flowID string) (bridgev2.LoginProcess, error) {
	return nil, fmt.Errorf("login not supported")
}

// JoinRoom joins the specified room
func (rc *RoomJoinerConnector) JoinRoom(roomID id.RoomID) error {
	fmt.Printf("Attempting to join room: %s\n", roomID.String())
	return nil
}

func (rc *RoomJoinerConnector) LoadUserLogin(ctx context.Context, login *bridgev2.UserLogin) error {
	return nil
}
