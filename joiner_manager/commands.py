# joiner_manager/commands.py
from mautrix.bridge.commands import CommandEvent, CommandHandler

class CommandHandler:
    def __init__(self, bridge: 'JoinerManagerBridge'):
        self.bridge = bridge

    async def handle_join(self, evt: CommandEvent):
        if len(evt.args) < 1:
            await evt.reply("Usage: !jm join <room_id>")
            return
        room_id = evt.args[0]
        await self.bridge.join_room(room_id)
        await evt.reply(f"Joined room {room_id}")

    async def handle_list(self, evt: CommandEvent):
        rooms = await self.bridge.db.fetch("SELECT room_id FROM managed_rooms")
        if not rooms:
            await evt.reply("No managed rooms")
            return
        response = "Managed rooms:\n" + "\n".join([f"- {room['room_id']}" for room in rooms])
        await evt.reply(response)