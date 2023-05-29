#include "steam_export.h"


void ReleaseSteamNetworkingMessageWrapper(SteamNetworkingMessage_t* message) {
    message->Release();
}
