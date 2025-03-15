# joiner_manager/__main__.py
from mautrix.bridge import Bridge
from . import JoinerManagerBridge

Bridge.run(JoinerManagerBridge)
