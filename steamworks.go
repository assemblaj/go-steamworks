// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2021 The go-steamworks Authors
//go:generate go run gen.go

package steamworks

type AppId_t uint32
type CSteamID uint64
type InputHandle_t uint64

type ESteamInputType int32
type EResult int32

const (
	EResultNone         EResult = 0
	EResultOK                   = 1
	EResultFail                 = 2
	EResultNoConnection         = 3
	// EResultNoConnectionRetry = 4				// OBSOLETE - removed
	EResultInvalidPassword                         = 5  // password/ticket is invalid
	EResultLoggedInElsewhere                       = 6  // same user logged in elsewhere
	EResultInvalidProtocolVer                      = 7  // protocol version is incorrect
	EResultInvalidParam                            = 8  // a parameter is incorrect
	EResultFileNotFound                            = 9  // file was not found
	EResultBusy                                    = 10 // called method busy - action not taken
	EResultInvalidState                            = 11 // called object was in an invalid state
	EResultInvalidName                             = 12 // name is invalid
	EResultInvalidEmail                            = 13 // email is invalid
	EResultDuplicateName                           = 14 // name is not unique
	EResultAccessDenied                            = 15 // access is denied
	EResultTimeout                                 = 16 // operation timed out
	EResultBanned                                  = 17 // VAC2 banned
	EResultAccountNotFound                         = 18 // account not found
	EResultInvalidSteamID                          = 19 // steamID is invalid
	EResultServiceUnavailable                      = 20 // The requested service is currently unavailable
	EResultNotLoggedOn                             = 21 // The user is not logged on
	EResultPending                                 = 22 // Request is pending (may be in process or waiting on third party)
	EResultEncryptionFailure                       = 23 // Encryption or Decryption failed
	EResultInsufficientPrivilege                   = 24 // Insufficient privilege
	EResultLimitExceeded                           = 25 // Too much of a good thing
	EResultRevoked                                 = 26 // Access has been revoked (used for revoked guest passes)
	EResultExpired                                 = 27 // License/Guest pass the user is trying to access is expired
	EResultAlreadyRedeemed                         = 28 // Guest pass has already been redeemed by account cannot be acked again
	EResultDuplicateRequest                        = 29 // The request is a duplicate and the action has already occurred in the past ignored this time
	EResultAlreadyOwned                            = 30 // All the games in this guest pass redemption request are already owned by the user
	EResultIPNotFound                              = 31 // IP address not found
	EResultPersistFailed                           = 32 // failed to write change to the data store
	EResultLockingFailed                           = 33 // failed to acquire access lock for this operation
	EResultLogonSessionReplaced                    = 34
	EResultConnectFailed                           = 35
	EResultHandshakeFailed                         = 36
	EResultIOFailure                               = 37
	EResultRemoteDisconnect                        = 38
	EResultShoppingCartNotFound                    = 39 // failed to find the shopping cart requested
	EResultBlocked                                 = 40 // a user didn't allow it
	EResultIgnored                                 = 41 // target is ignoring sender
	EResultNoMatch                                 = 42 // nothing matching the request found
	EResultAccountDisabled                         = 43
	EResultServiceReadOnly                         = 44 // this service is not accepting content changes right now
	EResultAccountNotFeatured                      = 45 // account doesn't have value so this feature isn't available
	EResultAdministratorOK                         = 46 // allowed to take this action but only because requester is admin
	EResultContentVersion                          = 47 // A Version mismatch in content transmitted within the Steam protocol.
	EResultTryAnotherCM                            = 48 // The current CM can't service the user making a request user should try another.
	EResultPasswordRequiredToKickSession           = 49 // You are already logged in elsewhere this cached credential login has failed.
	EResultAlreadyLoggedInElsewhere                = 50 // You are already logged in elsewhere you must wait
	EResultSuspended                               = 51 // Long running operation (content download) suspended/paused
	EResultCancelled                               = 52 // Operation canceled (typically by user: content download)
	EResultDataCorruption                          = 53 // Operation canceled because data is ill formed or unrecoverable
	EResultDiskFull                                = 54 // Operation canceled - not enough disk space.
	EResultRemoteCallFailed                        = 55 // an remote call or IPC call failed
	EResultPasswordUnset                           = 56 // Password could not be verified as it's unset server side
	EResultExternalAccountUnlinked                 = 57 // External account (PSN Facebook...) is not linked to a Steam account
	EResultPSNTicketInvalid                        = 58 // PSN ticket was invalid
	EResultExternalAccountAlreadyLinked            = 59 // External account (PSN Facebook...) is already linked to some other account must explicitly request to replace/delete the link first
	EResultRemoteFileConflict                      = 60 // The sync cannot resume due to a conflict between the local and remote files
	EResultIllegalPassword                         = 61 // The requested new password is not legal
	EResultSameAsPreviousValue                     = 62 // new value is the same as the old one ( secret question and answer )
	EResultAccountLogonDenied                      = 63 // account login denied due to 2nd factor authentication failure
	EResultCannotUseOldPassword                    = 64 // The requested new password is not legal
	EResultInvalidLoginAuthCode                    = 65 // account login denied due to auth code invalid
	EResultAccountLogonDeniedNoMail                = 66 // account login denied due to 2nd factor auth failure - and no mail has been sent - partner site specific
	EResultHardwareNotCapableOfIPT                 = 67 //
	EResultIPTInitError                            = 68 //
	EResultParentalControlRestricted               = 69 // operation failed due to parental control restrictions for current user
	EResultFacebookQueryError                      = 70 // Facebook query returned an error
	EResultExpiredLoginAuthCode                    = 71 // account login denied due to auth code expired
	EResultIPLoginRestrictionFailed                = 72
	EResultAccountLockedDown                       = 73
	EResultAccountLogonDeniedVerifiedEmailRequired = 74
	EResultNoMatchingURL                           = 75
	EResultBadResponse                             = 76  // parse failure missing field etc.
	EResultRequirePasswordReEntry                  = 77  // The user cannot complete the action until they re-enter their password
	EResultValueOutOfRange                         = 78  // the value entered is outside the acceptable range
	EResultUnexpectedError                         = 79  // something happened that we didn't expect to ever happen
	EResultDisabled                                = 80  // The requested service has been configured to be unavailable
	EResultInvalidCEGSubmission                    = 81  // The set of files submitted to the CEG server are not valid !
	EResultRestrictedDevice                        = 82  // The device being used is not allowed to perform this action
	EResultRegionLocked                            = 83  // The action could not be complete because it is region restricted
	EResultRateLimitExceeded                       = 84  // Temporary rate limit exceeded try again later different from EResultLimitExceeded which may be permanent
	EResultAccountLoginDeniedNeedTwoFactor         = 85  // Need two-factor code to login
	EResultItemDeleted                             = 86  // The thing we're trying to access has been deleted
	EResultAccountLoginDeniedThrottle              = 87  // login attempt failed try to throttle response to possible attacker
	EResultTwoFactorCodeMismatch                   = 88  // two factor code mismatch
	EResultTwoFactorActivationCodeMismatch         = 89  // activation code for two-factor didn't match
	EResultAccountAssociatedToMultiplePartners     = 90  // account has been associated with multiple partners
	EResultNotModified                             = 91  // data not modified
	EResultNoMobileDevice                          = 92  // the account does not have a mobile device associated with it
	EResultTimeNotSynced                           = 93  // the time presented is out of range or tolerance
	EResultSmsCodeFailed                           = 94  // SMS code failure (no match none pending etc.)
	EResultAccountLimitExceeded                    = 95  // Too many accounts access this resource
	EResultAccountActivityLimitExceeded            = 96  // Too many changes to this account
	EResultPhoneActivityLimitExceeded              = 97  // Too many changes to this phone
	EResultRefundToWallet                          = 98  // Cannot refund to payment method must use wallet
	EResultEmailSendFailure                        = 99  // Cannot send an email
	EResultNotSettled                              = 100 // Can't perform operation till payment has settled
	EResultNeedCaptcha                             = 101 // Needs to provide a valid captcha
	EResultGSLTDenied                              = 102 // a game server login token owned by this token's owner has been banned
	EResultGSOwnerDenied                           = 103 // game server owner is denied for other reason (account lock community ban vac ban missing phone)
	EResultInvalidItemType                         = 104 // the type of thing we were requested to act on is invalid
	EResultIPBanned                                = 105 // the ip address has been banned from taking this action
	EResultGSLTExpired                             = 106 // this token has expired from disuse; can be reset for use
	EResultInsufficientFunds                       = 107 // user doesn't have enough wallet funds to complete the action
	EResultTooManyPending                          = 108 // There are too many of this thing pending already
	EResultNoSiteLicensesFound                     = 109 // No site licenses found
	EResultWGNetworkSendExceeded                   = 110 // the WG couldn't send a response because we exceeded max network send size
	EResultAccountNotFriends                       = 111 // the user is not mutually friends
	EResultLimitedUserAccount                      = 112 // the user is limited
	EResultCantRemoveItem                          = 113 // item can't be removed
	EResultAccountDeleted                          = 114 // account has been deleted
	EResultExistingUserCancelledLicense            = 115 // A license for this already exists but cancelled
	EResultCommunityCooldown                       = 116 // access is denied because of a community cooldown (probably from support profile data resets)
	EResultNoLauncherSpecified                     = 117 // No launcher was specified but a launcher was needed to choose correct realm for operation.
	EResultMustAgreeToSSA                          = 118 // User must agree to china SSA or global SSA before login
	EResultLauncherMigrated                        = 119 // The specified launcher type is no longer supported; the user should be directed elsewhere
	EResultSteamRealmMismatch                      = 120 // The user's realm does not match the realm of the requested resource
	EResultInvalidSignature                        = 121 // signature check did not match
	EResultParseFailure                            = 122 // Failed to parse input
	EResultNoVerifiedPhone                         = 123 // account does not have a verified phone number
	EResultInsufficientBattery                     = 124 // user device doesn't have enough battery charge currently to complete the action
	EResultChargerRequired                         = 125 // The operation requires a charger to be plugged in which wasn't present
	EResultCachedCredentialInvalid                 = 126 // Cached credential was invalid - user must reauthenticate
	EResultPhoneNumberIsVOIP                       = 127 // The phone number provided is a Voice Over IP number
)

type HSteamListenSocket int32
type HSteamNetConnection int32
type ESteamNetworkingConnectionState int32
type SteamNetworkingPOPID uint32
type SteamNetworkingMicroseconds int64

const iSteamNetworkingSocketsCallbacks = 1220 // Assuming it's 1, modify as required.
const iCallback = iSteamNetworkingSocketsCallbacks + 1

type SteamNetConnectionStatusChangedCallback_t struct {
	Conn     HSteamNetConnection
	Info     SteamNetConnectionInfo_t
	OldState ESteamNetworkingConnectionState
}

const (
	ESteamNetworkingConnectionState_None                   ESteamNetworkingConnectionState = 0
	ESteamNetworkingConnectionState_Connecting                                             = 1
	ESteamNetworkingConnectionState_FindingRoute                                           = 2
	ESteamNetworkingConnectionState_Connected                                              = 3
	ESteamNetworkingConnectionState_ClosedByPeer                                           = 4
	ESteamNetworkingConnectionState_ProblemDetectedLocally                                 = 5
	ESteamNetworkingConnectionState_FinWait                                                = -1
	ESteamNetworkingConnectionState_Linger                                                 = -2
	ESteamNetworkingConnectionState_Dead                                                   = -3
	ESteamNetworkingConnectionState__Force32Bit                                            = 0x7fffffff
)

const (
	cchSteamNetworkingMaxConnectionCloseReason = 128
	cchSteamNetworkingMaxConnectionDescription = 128
)

const (
	nSteamNetworkConnectionInfoFlags_Unauthenticated = 1
	nSteamNetworkConnectionInfoFlags_Unencrypted     = 2
	nSteamNetworkConnectionInfoFlags_LoopbackBuffers = 4
	nSteamNetworkConnectionInfoFlags_Fast            = 8
	nSteamNetworkConnectionInfoFlags_Relayed         = 16
	nSteamNetworkConnectionInfoFlags_DualWifi        = 32
)

type IPv4MappedAddress struct {
	EightZeros uint64
	Zeros      uint16
	FFFF       uint16
	IP         [4]byte // In network byte order
}

type SteamNetworkingIPAddr struct {
	IPv6 [16]byte
	IPv4 IPv4MappedAddress
	Port uint16
}

type ESteamNetworkingIdentityType int32

type ESteamNetworkingFakeIPType int32

type SteamNetworkingIdentity struct {
	EType            ESteamNetworkingIdentityType
	SteamID64        uint64
	PSNID            uint64
	StadiaID         uint64
	GenericString    [32]byte
	XboxPairwiseID   [33]byte
	GenericBytes     [32]byte
	UnknownRawString [128]byte
	IP               SteamNetworkingIPAddr
	Reserved         [32]uint32
}

type SteamNetConnectionInfo_t struct {
	identityRemote        SteamNetworkingIdentity
	UserData              uint64
	ListenSocket          HSteamListenSocket
	addrRemote            SteamNetworkingIPAddr
	pad1                  uint16
	idPOPRemote           SteamNetworkingPOPID
	idPOPRelay            SteamNetworkingPOPID
	eState                ESteamNetworkingConnectionState
	EndReason             int
	EndDebug              [cchSteamNetworkingMaxConnectionCloseReason]byte
	ConnectionDescription [cchSteamNetworkingMaxConnectionDescription]byte
	Flags                 int /// Misc flags.  Bitmask of nSteamNetworkConnectionInfoFlags_Xxxx
	reserved              [63]uint32
}

type SteamNetConnectionRealTimeStatus_t struct {
	/// High level state of the connection
	m_eState ESteamNetworkingConnectionState
	/// Current ping (ms)
	m_nPing int

	/// Connection quality measured locally, 0...1.  (Percentage of packets delivered
	/// end-to-end in order).
	m_flConnectionQualityLocal float64

	/// Packet delivery success rate as observed from remote host
	m_flConnectionQualityRemote float64

	/// Current data rates from recent history.
	m_flOutPacketsPerSec float64
	m_flOutBytesPerSec   float64
	m_flInPacketsPerSec  float64
	m_flInBytesPerSec    float64

	m_nSendRateBytesPerSecond int
	m_cbPendingUnreliable     int
	m_cbPendingReliable       int
	/// Number of bytes of reliable data that has been placed the wire, but
	/// for which we have not yet received an acknowledgment, and thus we may
	/// have to re-transmit.
	m_cbSentUnackedReliable int
	m_usecQueueTime         SteamNetworkingMicroseconds
	// Internal stuff, room to change API easily
	reserved [16]uint32
}

type SteamNetworkingMessage_t struct {
	Data          *byte // m_pData
	Size          int   // m_cbSize
	Connection    HSteamNetConnection
	PeerIdentity  SteamNetworkingIdentity // m_identityPeer
	ConnUserData  int64
	TimeReceived  SteamNetworkingMicroseconds // m_usecTimeReceived
	MessageNumber int64                       // m_nMessageNumber
	Channel       int                         // m_nChannel
	Flags         int                         // m_nFlags
	UserData      int64                       // m_nUserData
	LaneIdx       uint16                      // m_idxLane
	Padding       uint16                      // _pad1__
}

type ELobbyType int

const (
	k_ELobbyTypePrivate     ELobbyType = 0 // only way to join the lobby is to invite to someone else
	k_ELobbyTypeFriendsOnly ELobbyType = 1 // shows for friends or invitees, but not in lobby list
	k_ELobbyTypePublic      ELobbyType = 2 // visible for friends and in lobby list
	k_ELobbyTypeInvisible   ELobbyType = 3 // returned by search, but not visible to other friends
	//    useful if you want a user in two lobbies, for example matching groups together
	//	  a user can be in only one regular lobby, and up to two invisible lobbies
	k_ELobbyTypePrivateUnique ELobbyType = 4 // private, unique and does not delete when empty - only one of these may exist per unique keypair set
)

type LobbyCreated_t struct {
	m_eResult EResult
	// k_EResultOK - the lobby was successfully created
	// k_EResultNoConnection - your Steam client doesn't have a connection to the back-end
	// k_EResultTimeout - you the message to the Steam servers, but it didn't respond
	// k_EResultFail - the server responded, but with an unknown internal error
	// k_EResultAccessDenied - your game isn't set to allow lobbies, or your client does haven't rights to play the game
	// k_EResultLimitExceeded - your game client has created too many lobbies

	m_ulSteamIDLobby uint64 // chat room, zero if failed
}

type LobbyMatchList_t struct {
	m_nLobbiesMatching uint32
}

type LobbyEnter_t struct {
	m_ulSteamIDLobby         uint64 // SteamID of the Lobby you have entered
	m_rgfChatPermissions     uint32 // Permissions of the current user
	m_bLocked                bool   // If true, then only invited users may join
	m_EChatRoomEnterResponse uint32 // EChatRoomEnterResponse
}

type SteamAPICallbackHandle uint64

type ELobbyComparison int
type ELobbyDistanceFilter int
type EChatEntryType int

type HSteamPipe int32
type HSteamUser int32

type CallbackMsg_t struct {
	m_hSteamUser HSteamUser // Specific user to whom this callback applies.
	m_iCallback  int        // Callback identifier.  (Corresponds to the k_iCallback enum in the callback structure.)
	m_pubParam   *uint8     // Points to the callback structure
	m_cubParam   int        // Size of the data pointed to by m_pubParam
}

type SteamCallbackID uint32

const (
	k_iSteamUserCallbacks               SteamCallbackID = 100
	k_iSteamGameServerCallbacks                         = 200
	k_iSteamFriendsCallbacks                            = 300
	k_iSteamBillingCallbacks                            = 400
	k_iSteamMatchmakingCallbacks                        = 500
	k_iSteamContentServerCallbacks                      = 600
	k_iSteamUtilsCallbacks                              = 700
	k_iSteamAppsCallbacks                               = 1000
	k_iSteamUserStatsCallbacks                          = 1100
	k_iSteamNetworkingCallbacks                         = 1200
	k_iSteamNetworkingSocketsCallbacks                  = 1220
	k_iSteamNetworkingMessagesCallbacks                 = 1250
	k_iSteamNetworkingUtilsCallbacks                    = 1280
	k_iSteamRemoteStorageCallbacks                      = 1300
	k_iSteamGameServerItemsCallbacks                    = 1500
	k_iSteamGameCoordinatorCallbacks                    = 1700
	k_iSteamGameServerStatsCallbacks                    = 1800
	k_iSteam2AsyncCallbacks                             = 1900
	k_iSteamGameStatsCallbacks                          = 2000
	k_iSteamHTTPCallbacks                               = 2100
	k_iSteamScreenshotsCallbacks                        = 2300
	k_iSteamStreamLauncherCallbacks                     = 2600
	k_iSteamControllerCallbacks                         = 2800
	k_iSteamUGCCallbacks                                = 3400
	k_iSteamStreamClientCallbacks                       = 3500
	k_iSteamAppListCallbacks                            = 3900
	k_iSteamMusicCallbacks                              = 4000
	k_iSteamMusicRemoteCallbacks                        = 4100
	k_iSteamGameNotificationCallbacks                   = 4400
	k_iSteamHTMLSurfaceCallbacks                        = 4500
	k_iSteamVideoCallbacks                              = 4600
	k_iSteamInventoryCallbacks                          = 4700
	k_ISteamParentalSettingsCallbacks                   = 5000
	k_iSteamGameSearchCallbacks                         = 5200
	k_iSteamPartiesCallbacks                            = 5300
	k_iSteamSTARCallbacks                               = 5500
	k_iSteamRemotePlayCallbacks                         = 5700
	k_iSteamChatCallbacks                               = 5900
)

type SteamAPICallCompleted_t struct {
	m_hAsyncCall SteamAPICallbackHandle
	m_iCallback  int
	m_cubParam   uint32
}

const (
	ESteamInputType_Unknown              ESteamInputType = 0
	ESteamInputType_SteamController      ESteamInputType = 1
	ESteamInputType_XBox360Controller    ESteamInputType = 2
	ESteamInputType_XBoxOneController    ESteamInputType = 3
	ESteamInputType_GenericXInput        ESteamInputType = 4
	ESteamInputType_PS4Controller        ESteamInputType = 5
	ESteamInputType_AppleMFiController   ESteamInputType = 6 // Unused
	ESteamInputType_AndroidController    ESteamInputType = 7 // Unused
	ESteamInputType_SwitchJoyConPair     ESteamInputType = 8 // Unused
	ESteamInputType_SwitchJoyConSingle   ESteamInputType = 9 // Unused
	ESteamInputType_SwitchProController  ESteamInputType = 10
	ESteamInputType_MobileTouch          ESteamInputType = 11
	ESteamInputType_PS3Controller        ESteamInputType = 12
	ESteamInputType_PS5Controller        ESteamInputType = 13
	ESteamInputType_SteamDeckController  ESteamInputType = 14
	ESteamInputType_Count                ESteamInputType = 15
	ESteamInputType_MaximumPossibleValue ESteamInputType = 255
)

const (
	_STEAM_INPUT_MAX_COUNT = 16
)

type ISteamApps interface {
	GetAppInstallDir(appID AppId_t) string
	GetCurrentGameLanguage() string
}

type ISteamInput interface {
	GetConnectedControllers() []InputHandle_t
	GetInputTypeForHandle(inputHandle InputHandle_t) ESteamInputType
	Init(bExplicitlyCallRunFrame bool) bool
	RunFrame()
}

type ISteamRemoteStorage interface {
	FileWrite(file string, data []byte) bool
	FileRead(file string, data []byte) int32
	FileDelete(file string) bool
	GetFileSize(file string) int32
}

type ISteamUser interface {
	GetSteamID() CSteamID
}

type ISteamUserStats interface {
	RequestCurrentStats() bool
	GetAchievement(name string) (achieved, success bool)
	SetAchievement(name string) bool
	ClearAchievement(name string) bool
	StoreStats() bool
}

type ISteamUtils interface {
	IsSteamRunningOnSteamDeck() bool
}

// type ISteamNetworkingSockets interface {
// 	CreateListenSocketIP(ip net.IP, port int) HSteamListenSocket
// 	ConnectByIPAddress(ip net.IP, port int) HSteamNetConnection
// 	AcceptConnection(connection HSteamNetConnection) EResult
// 	CloseConnection(connection HSteamNetConnection, reason int, debug string, userHasSentDisconnect bool)
// 	CloseListenSocket(socket HSteamListenSocket)
// 	RunCallbacks()
// }

type ISteamNetworkingMessages interface {
	SendMessageToUser(identity SteamNetworkingIdentity, data []byte, sendFlags int, channel int) EResult
	ReceiveMessagesOnChannel(localChannel int, maxMessages int) ([]SteamNetworkingMessage_t, EResult)
	AcceptSessionWithUser(identityRemote SteamNetworkingIdentity) bool
	CloseSessionWithUser(identityRemote SteamNetworkingIdentity) bool
	CloseChannelWithUser(identityRemote SteamNetworkingIdentity, nLocalChannel int) bool
	GetSessionConnectionInfo(identityRemote SteamNetworkingIdentity) (ESteamNetworkingConnectionState, SteamNetConnectionInfo_t, SteamNetConnectionRealTimeStatus_t)
}

type ISteamMatchmaking interface {
	CreateLobby(eLobbyType ELobbyType, cMaxMembers int) (LobbyCreated_t, error)
	RequestLobbyList() (list LobbyMatchList_t, err error)
	// GetLobbyByIndex(iLobby int) CSteamID
	// JoinLobby(steamIDLobby CSteamID) SteamAPICall_t
	GetLobbyByIndex(iLobby int) CSteamID
	LeaveLobby(steamIDLobby CSteamID)
	// InviteUserToLobby(steamIDLobby, steamIDInvitee CSteamID) bool
	// GetNumLobbyMembers(steamIDLobby CSteamID) int

	// GetLobbyData(steamIDLobby CSteamID, pchKey string) string
	// SetLobbyData(steamIDLobby CSteamID, pchKey, pchValue string) bool
	// GetLobbyDataCount(steamIDLobby CSteamID) int
	// GetLobbyDataByIndex(steamIDLobby CSteamID, iLobbyData int, cchKeyBufferSize int, cchValueBufferSize int) (bool, []string, []string)
	//DeleteLobbyData(steamIDLobby CSteamID, pchKey string) bool
	// GetLobbyMemberData(steamIDLobby CSteamID, steamIDUser, pchKey string) string
	// SetLobbyMemberData(steamIDLobby CSteamID, pchKey, pchValue string)
	// SendLobbyChatMsg(steamIDLobby CSteamID, pvMsgBody []byte, cubMsgBody int) bool
	// GetLobbyChatEntry(steamIDLobby CSteamID, iChatID int, pSteamIDUser string, pvData []byte, cubData int, peChatEntryType EChatEntryType) int
	// RequestLobbyData(steamIDLobby CSteamID) bool
	// SetLobbyGameServer(steamIDLobby CSteamID, unGameServerIP uint32, unGameServerPort uint16, steamIDGameServer string)
	// GetLobbyGameServer(steamIDLobby CSteamID, punGameServerIP *uint32, punGameServerPort *uint16, psteamIDGameServer *string) bool
	// SetLobbyMemberLimit(steamIDLobby CSteamID, cMaxMembers int) bool
	// GetLobbyMemberLimit(steamIDLobby CSteamID) int
	// SetLobbyType(steamIDLobby CSteamID, eLobbyType ELobbyType) bool
	// SetLobbyJoinable(steamIDLobby CSteamID, bLobbyJoinable bool) bool
	// GetLobbyOwner(steamIDLobby CSteamID) CSteamID
	// SetLobbyOwner(steamIDLobby, steamIDNewOwner CSteamID) bool
	// SetLinkedLobby(steamIDLobby, steamIDLobbyDependent CSteamID) bool
}

const (
	flatAPI_RestartAppIfNecessary = "SteamAPI_RestartAppIfNecessary"
	flatAPI_Init                  = "SteamAPI_Init"
	flatAPI_RunCallbacks          = "SteamAPI_RunCallbacks"

	flatAPI_SteamApps                         = "SteamAPI_SteamApps_v008"
	flatAPI_ISteamApps_GetAppInstallDir       = "SteamAPI_ISteamApps_GetAppInstallDir"
	flatAPI_ISteamApps_GetCurrentGameLanguage = "SteamAPI_ISteamApps_GetCurrentGameLanguage"

	flatAPI_SteamInput                          = "SteamAPI_SteamInput_v006"
	flatAPI_ISteamInput_GetConnectedControllers = "SteamAPI_ISteamInput_GetConnectedControllers"
	flatAPI_ISteamInput_GetInputTypeForHandle   = "SteamAPI_ISteamInput_GetInputTypeForHandle"
	flatAPI_ISteamInput_Init                    = "SteamAPI_ISteamInput_Init"
	flatAPI_ISteamInput_RunFrame                = "SteamAPI_ISteamInput_RunFrame"

	flatAPI_SteamRemoteStorage              = "SteamAPI_SteamRemoteStorage_v016"
	flatAPI_ISteamRemoteStorage_FileWrite   = "SteamAPI_ISteamRemoteStorage_FileWrite"
	flatAPI_ISteamRemoteStorage_FileRead    = "SteamAPI_ISteamRemoteStorage_FileRead"
	flatAPI_ISteamRemoteStorage_FileDelete  = "SteamAPI_ISteamRemoteStorage_FileDelete"
	flatAPI_ISteamRemoteStorage_GetFileSize = "SteamAPI_ISteamRemoteStorage_GetFileSize"

	flatAPI_SteamUser             = "SteamAPI_SteamUser_v021"
	flatAPI_ISteamUser_GetSteamID = "SteamAPI_ISteamUser_GetSteamID"

	flatAPI_SteamUserStats                      = "SteamAPI_SteamUserStats_v012"
	flatAPI_ISteamUserStats_RequestCurrentStats = "SteamAPI_ISteamUserStats_RequestCurrentStats"
	flatAPI_ISteamUserStats_GetAchievement      = "SteamAPI_ISteamUserStats_GetAchievement"
	flatAPI_ISteamUserStats_SetAchievement      = "SteamAPI_ISteamUserStats_SetAchievement"
	flatAPI_ISteamUserStats_ClearAchievement    = "SteamAPI_ISteamUserStats_ClearAchievement"
	flatAPI_ISteamUserStats_StoreStats          = "SteamAPI_ISteamUserStats_StoreStats"

	flatAPI_SteamUtils                            = "SteamAPI_SteamUtils_v010"
	flatAPI_ISteamUtils_IsSteamRunningOnSteamDeck = "SteamAPI_ISteamUtils_IsSteamRunningOnSteamDeck"

	flatAPI_SteamNetworkingMessages                           = "SteamAPI_SteamNetworkingMessages_SteamAPI_v002"
	flatAPI_ISteamNetworkingMessages_SendMessageToUser        = "SteamAPI_ISteamNetworkingMessages_SendMessageToUser"
	flatAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel = "SteamAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel "
	flatAPI_ISteamNetworkingMessages_AcceptSessionWithUser    = "SteamAPI_ISteamNetworkingMessages_AcceptSessionWithUser"
	flatAPI_ISteamNetworkingMessages_CloseSessionWithUser     = "SteamAPI_ISteamNetworkingMessages_CloseSessionWithUser"
	flatAPI_ISteamNetworkingMessages_CloseChannelWithUser     = "SteamAPI_ISteamNetworkingMessages_CloseChannelWithUser"
	flatAPI_ISteamNetworkingMessages_GetSessionConnectionInfo = "SteamAPI_ISteamNetworkingMessages_GetSessionConnectionInfo"

	flatAPI_SteamAPI_SteamMatchmaking                   = "SteamAPI_SteamMatchmaking_v009"
	flatAPI_SteamAPI_ISteamMatchmaking_RequestLobbyList = "SteamAPI_ISteamMatchmaking_RequestLobbyList"
	flatAPI_SteamAPI_ISteamMatchmaking_GetLobbyByIndex  = "SteamAPI_ISteamMatchmaking_GetLobbyByIndex"
	flatAPI_SteamAPI_ISteamMatchmaking_CreateLobby      = "SteamAPI_ISteamMatchmaking_CreateLobby"
	flatAPI_SteamAPI_ISteamMatchmaking_LeaveLobby       = "SteamAPI_ISteamMatchmaking_LeaveLobby"

	SteamAPI_GetHSteamPipe                   = "SteamAPI_GetHSteamPipe"
	SteamAPI_ManualDispatch_Init             = "SteamAPI_ManualDispatch_Init"
	SteamAPI_ManualDispatch_RunFrame         = "SteamAPI_ManualDispatch_RunFrame"
	SteamAPI_ManualDispatch_GetNextCallback  = "SteamAPI_ManualDispatch_GetNextCallback"
	SteamAPI_ManualDispatch_FreeLastCallback = "SteamAPI_ManualDispatch_FreeLastCallback"
	SteamAPI_ManualDispatch_GetAPICallResult = "SteamAPI_ManualDispatch_GetAPICallResult"
)
