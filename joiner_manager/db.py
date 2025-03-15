# joiner_manager/db.py
from mautrix.util.async_db import Database, Upgrade

upgrade_table = [
    Upgrade(
        "Initial schema",
        """CREATE TABLE IF NOT EXISTS managed_rooms (
            room_id TEXT PRIMARY KEY,
            joined_at BIGINT NOT NULL
        )""",
    ),
]