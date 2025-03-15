# Joiner Bridge

A simple Matrix bridge that joins rooms.

## Setup

1. Copy `example-config.yaml` to `config.yaml` and edit it with your homeserver details
2. Run the bridge once to generate the registration file:
   ```bash
   ./joiner-bridge -g
   ```
3. Register the bridge with your homeserver by copying the generated `registration.yaml` file to your homeserver's appservice directory
4. Start the bridge:
   ```bash
   ./joiner-bridge
   ```

## Usage

1. Invite the bridge bot (@joinerbot:yourdomain) to a room
2. The bot will automatically join the room
3. You can use the following commands in rooms where the bot is present:
   - `!joiner help` - Show help message
   - `!joiner join <room_id>` - Join a room by ID
   - `!joiner leave <room_id>` - Leave a room by ID

## Building from Source

```bash
go build ./cmd/joiner-bridge
```
