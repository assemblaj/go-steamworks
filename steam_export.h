#ifndef STEAM_EXPORT_H
#define STEAM_EXPORT_H

#include <steam/steam_api.h>

#ifdef __cplusplus
extern "C" {
#endif

void ReleaseSteamNetworkingMessageWrapper(SteamNetworkingMessage_t* message);

#ifdef __cplusplus
}
#endif

#endif /* STEAM_EXPORT_H */
