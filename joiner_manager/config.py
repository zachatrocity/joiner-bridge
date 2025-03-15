# joiner_manager/config.py
from mautrix.util.config import BaseProxyConfig, ConfigUpdateHelper

class Config(BaseProxyConfig):
    def do_update(self, helper: ConfigUpdateHelper) -> None:
        helper.copy("homeserver.address")
        helper.copy("homeserver.domain")
        helper.copy("appservice.database")
        helper.copy("appservice.hostname")
        helper.copy("appservice.port")
        helper.copy("appservice.max_retries")
        helper.copy("bridge.permissions")