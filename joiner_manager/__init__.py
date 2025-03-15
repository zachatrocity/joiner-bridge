# joiner_manager/__init__.py
from mautrix.bridge import Bridge
from mautrix.types import RoomID
from mautrix.util.async_db import Database
import time

class JoinerManagerBridge(Bridge):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.db = Database()

    @classmethod
    def get_config_class(cls):
        from .config import Config
        return Config

    @classmethod
    def get_db_upgrade_table(cls):
        return "joinermanager_version"

    @classmethod
    def get_db_class(cls):
        return Database

    async def start(self) -> None:
        await super().start()
        self.config.load_and_update()
        await self.db.start()

    async def join_room(self, room_id: RoomID) -> None:
        try:
            await self.az.intent.join_room(room_id)
            await self.db.execute(
                "INSERT INTO managed_rooms (room_id, joined_at) VALUES (?, ?)",
                room_id, int(time.time())
            )
        except Exception as e:
            self.log.error(f"Failed to join room {room_id}: {e}")