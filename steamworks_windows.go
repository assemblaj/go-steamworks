// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2021 The go-steamworks Authors

package steamworks

import (
	"errors"
	"runtime"
	"unsafe"

	"golang.org/x/sys/windows"
)

const is32Bit = unsafe.Sizeof(int(0)) == 4

type dll struct {
	d     *windows.LazyDLL
	procs map[string]*windows.LazyProc
}

func (d *dll) call(name string, args ...uintptr) (uintptr, error) {
	if d.procs == nil {
		d.procs = map[string]*windows.LazyProc{}
	}
	if _, ok := d.procs[name]; !ok {
		d.procs[name] = d.d.NewProc(name)
	}
	r, _, err := d.procs[name].Call(args...)
	if err != nil {
		errno, ok := err.(windows.Errno)
		if !ok {
			return r, err
		}
		if errno != 0 {
			return r, err
		}
	}
	return r, nil
}

func loadDLL() (*dll, error) {
	dllName := "steam_api.dll"
	if !is32Bit {
		dllName = "steam_api64.dll"
	}

	return &dll{
		d: windows.NewLazyDLL(dllName),
	}, nil
}

var theDLL *dll

func init() {
	dll, err := loadDLL()
	if err != nil {
		panic(err)
	}
	theDLL = dll
}

func RestartAppIfNecessary(appID uint32) bool {
	v, err := theDLL.call(flatAPI_RestartAppIfNecessary, uintptr(appID))
	if err != nil {
		panic(err)
	}
	return byte(v) != 0
}

func Init() bool {
	v, err := theDLL.call(flatAPI_Init)
	if err != nil {
		panic(err)
	}
	manualDispatch_Init()
	return byte(v) != 0
}

func RunCallbacks() {
	if _, err := theDLL.call(flatAPI_RunCallbacks); err != nil {
		panic(err)
	}
}

func SteamApps() ISteamApps {
	v, err := theDLL.call(flatAPI_SteamApps)
	if err != nil {
		panic(err)
	}
	return steamApps(v)
}

type steamApps uintptr

func (s steamApps) GetAppInstallDir(appID AppId_t) string {
	var path [4096]byte
	v, err := theDLL.call(flatAPI_ISteamApps_GetAppInstallDir, uintptr(s), uintptr(appID), uintptr(unsafe.Pointer(&path[0])), uintptr(len(path)))
	if err != nil {
		panic(err)
	}
	return string(path[:uint32(v)-1])
}

func (s steamApps) GetCurrentGameLanguage() string {
	v, err := theDLL.call(flatAPI_ISteamApps_GetCurrentGameLanguage, uintptr(s))
	if err != nil {
		panic(err)
	}

	bs := make([]byte, 0, 256)
	for i := int32(0); ; i++ {
		b := *(*byte)(unsafe.Pointer(v))
		v += unsafe.Sizeof(byte(0))
		if b == 0 {
			break
		}
		bs = append(bs, b)
	}
	return string(bs)
}

func SteamInput() ISteamInput {
	v, err := theDLL.call(flatAPI_SteamInput)
	if err != nil {
		panic(err)
	}
	return steamInput(v)
}

type steamInput uintptr

func (s steamInput) GetConnectedControllers() []InputHandle_t {
	var handles [_STEAM_INPUT_MAX_COUNT]InputHandle_t
	v, err := theDLL.call(flatAPI_ISteamInput_GetConnectedControllers, uintptr(s), uintptr(unsafe.Pointer(&handles[0])))
	if err != nil {
		panic(err)
	}
	return handles[:int(v)]
}

func (s steamInput) GetInputTypeForHandle(inputHandle InputHandle_t) ESteamInputType {
	v, err := theDLL.call(flatAPI_ISteamInput_GetInputTypeForHandle, uintptr(s), uintptr(inputHandle))
	if err != nil {
		panic(err)
	}
	return ESteamInputType(v)
}

func (s steamInput) Init(bExplicitlyCallRunFrame bool) bool {
	var callRunFrame uintptr
	if bExplicitlyCallRunFrame {
		callRunFrame = 1
	}
	// The error value seems unreliable.
	v, _ := theDLL.call(flatAPI_ISteamInput_Init, uintptr(s), callRunFrame)
	return byte(v) != 0
}

func (s steamInput) RunFrame() {
	if _, err := theDLL.call(flatAPI_ISteamInput_RunFrame, uintptr(s), 0); err != nil {
		panic(err)
	}
}

func SteamRemoteStorage() ISteamRemoteStorage {
	v, err := theDLL.call(flatAPI_SteamRemoteStorage)
	if err != nil {
		panic(err)
	}
	return steamRemoteStorage(v)
}

type steamRemoteStorage uintptr

func (s steamRemoteStorage) FileWrite(file string, data []byte) bool {
	cfile := append([]byte(file), 0)
	defer runtime.KeepAlive(cfile)

	defer runtime.KeepAlive(data)

	v, err := theDLL.call(flatAPI_ISteamRemoteStorage_FileWrite, uintptr(s), uintptr(unsafe.Pointer(&cfile[0])), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	if err != nil {
		panic(err)
	}

	return byte(v) != 0
}

func (s steamRemoteStorage) FileRead(file string, data []byte) int32 {
	cfile := append([]byte(file), 0)
	defer runtime.KeepAlive(cfile)

	defer runtime.KeepAlive(data)

	v, err := theDLL.call(flatAPI_ISteamRemoteStorage_FileRead, uintptr(s), uintptr(unsafe.Pointer(&cfile[0])), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	if err != nil {
		panic(err)
	}

	return int32(v)
}

func (s steamRemoteStorage) FileDelete(file string) bool {
	cfile := append([]byte(file), 0)
	defer runtime.KeepAlive(cfile)

	v, err := theDLL.call(flatAPI_ISteamRemoteStorage_FileDelete, uintptr(s), uintptr(unsafe.Pointer(&cfile[0])))
	if err != nil {
		panic(err)
	}

	return byte(v) != 0
}

func (s steamRemoteStorage) GetFileSize(file string) int32 {
	cfile := append([]byte(file), 0)
	defer runtime.KeepAlive(cfile)

	v, err := theDLL.call(flatAPI_ISteamRemoteStorage_GetFileSize, uintptr(s), uintptr(unsafe.Pointer(&cfile[0])))
	if err != nil {
		panic(err)
	}

	return int32(v)
}

func SteamUser() ISteamUser {
	v, err := theDLL.call(flatAPI_SteamUser)
	if err != nil {
		panic(err)
	}
	return steamUser(v)
}

type steamUser uintptr

func (s steamUser) GetSteamID() CSteamID {
	if is32Bit {
		// On 32bit machines, syscall cannot treat a returned value as 64bit.
		panic("GetSteamID is not implemented on 32bit Windows")
	}
	v, err := theDLL.call(flatAPI_ISteamUser_GetSteamID, uintptr(s))
	if err != nil {
		panic(err)
	}
	return CSteamID(v)
}

func SteamUserStats() ISteamUserStats {
	v, err := theDLL.call(flatAPI_SteamUserStats)
	if err != nil {
		panic(err)
	}
	return steamUserStats(v)
}

type steamUserStats uintptr

func (s steamUserStats) RequestCurrentStats() bool {
	v, err := theDLL.call(flatAPI_ISteamUserStats_RequestCurrentStats, uintptr(s))
	if err != nil {
		panic(err)
	}

	return byte(v) != 0
}

func (s steamUserStats) GetAchievement(name string) (achieved, success bool) {
	cname := append([]byte(name), 0)
	defer runtime.KeepAlive(cname)

	v, err := theDLL.call(flatAPI_ISteamUserStats_SetAchievement, uintptr(s), uintptr(unsafe.Pointer(&cname[0])), uintptr(unsafe.Pointer(&achieved)))
	if err != nil {
		panic(err)
	}

	success = byte(v) != 0
	return
}

func (s steamUserStats) SetAchievement(name string) bool {
	cname := append([]byte(name), 0)
	defer runtime.KeepAlive(cname)

	v, err := theDLL.call(flatAPI_ISteamUserStats_SetAchievement, uintptr(s), uintptr(unsafe.Pointer(&cname[0])))
	if err != nil {
		panic(err)
	}

	return byte(v) != 0
}

func (s steamUserStats) ClearAchievement(name string) bool {
	cname := append([]byte(name), 0)
	defer runtime.KeepAlive(cname)

	v, err := theDLL.call(flatAPI_ISteamUserStats_ClearAchievement, uintptr(s), uintptr(unsafe.Pointer(&cname[0])))
	if err != nil {
		panic(err)
	}

	return byte(v) != 0
}

func (s steamUserStats) StoreStats() bool {
	v, err := theDLL.call(flatAPI_ISteamUserStats_StoreStats, uintptr(s))
	if err != nil {
		panic(err)
	}

	return byte(v) != 0
}

func SteamUtils() ISteamUtils {
	v, err := theDLL.call(flatAPI_SteamUtils)
	if err != nil {
		panic(err)
	}
	return steamUtils(v)
}

type steamUtils uintptr

func (s steamUtils) IsSteamRunningOnSteamDeck() bool {
	v, err := theDLL.call(flatAPI_ISteamUtils_IsSteamRunningOnSteamDeck, uintptr(s))
	if err != nil {
		panic(err)
	}

	return byte(v) != 0
}

type steamNetworkingMessages uintptr

func SteamNetworkingMessages() ISteamNetworkingMessages {
	v, err := theDLL.call(flatAPI_SteamNetworkingMessages)
	if err != nil {
		panic(err)
	}
	return steamNetworkingMessages(v)
}

func (s steamNetworkingMessages) SendMessageToUser(identity SteamNetworkingIdentity, data []byte, sendFlags int32, channel int32) EResult {
	cID := uintptr(unsafe.Pointer(&identity))
	cData := uintptr(unsafe.Pointer(&data[0]))
	cLen := uintptr(len(data))
	cSendFlags := uintptr(sendFlags)
	cChannel := uintptr(channel)

	v, err := theDLL.call(flatAPI_ISteamNetworkingMessages_SendMessageToUser, uintptr(s), cID, cData, cLen, cSendFlags, cChannel)
	if err != nil {
		panic(err)
	}
	return EResult(v)
}

func (s steamNetworkingMessages) ReceiveMessagesOnChannel(localChannel int32, maxMessages int32) ([]SteamNetworkingMessage_t, EResult) {
	cLocalChannel := uintptr(localChannel)
	cMaxMessages := uintptr(maxMessages)
	data := make([]SteamNetworkingMessage_t, maxMessages)
	cData := uintptr(unsafe.Pointer(&data[0]))

	v, err := theDLL.call(flatAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel, uintptr(s), cLocalChannel, cData, cMaxMessages)
	if err != nil {
		panic(err)
	}
	result := EResult(v)
	switch result {
	case EResultOK:
		return data, result
	default:
		return nil, result
	}
}

func (s steamNetworkingMessages) AcceptSessionWithUser(identityRemote SteamNetworkingIdentity) bool {
	v, err := theDLL.call(flatAPI_ISteamNetworkingMessages_AcceptSessionWithUser, uintptr(s), uintptr(unsafe.Pointer(&identityRemote)))
	if err != nil {
		panic(err)
	}
	return byte(v) != 0
}

func (s steamNetworkingMessages) CloseSessionWithUser(identityRemote SteamNetworkingIdentity) bool {
	v, err := theDLL.call(flatAPI_ISteamNetworkingMessages_CloseSessionWithUser, uintptr(s), uintptr(unsafe.Pointer(&identityRemote)))
	if err != nil {
		panic(err)
	}
	return byte(v) != 0

}

func (s steamNetworkingMessages) CloseChannelWithUser(identityRemote SteamNetworkingIdentity, nLocalChannel int32) bool {
	v, err := theDLL.call(flatAPI_ISteamNetworkingMessages_CloseChannelWithUser, uintptr(s), uintptr(unsafe.Pointer(&identityRemote)), uintptr(nLocalChannel))
	if err != nil {
		panic(err)
	}
	return byte(v) != 0
}

func (s steamNetworkingMessages) GetSessionConnectionInfo(identityRemote SteamNetworkingIdentity) (ESteamNetworkingConnectionState, SteamNetConnectionInfo_t, SteamNetConnectionRealTimeStatus_t) {
	info := SteamNetConnectionInfo_t{}
	stats := SteamNetConnectionRealTimeStatus_t{}

	v, err := theDLL.call(flatAPI_ISteamNetworkingMessages_GetSessionConnectionInfo, uintptr(s), uintptr(unsafe.Pointer(&identityRemote)), uintptr(unsafe.Pointer(&info)), uintptr(unsafe.Pointer(&stats)))
	if err != nil {
		panic(err)
	}
	return ESteamNetworkingConnectionState(v), info, stats
}

func ReleaseMessages(messages []SteamNetworkingMessage_t) {
	for i := 0; i < len(messages); i++ {
		_, err := theDLL.call(flatAPI_SteamAPI_SteamNetworkingMessage_t_Release, uintptr(unsafe.Pointer(&messages[i])))
		if err != nil {
			panic(err)
		}
	}
}

type steamMatchmaking uintptr

func SteamMatchmaking() ISteamMatchmaking {
	v, err := theDLL.call(flatAPI_SteamMatchmaking)
	if err != nil {
		panic(err)
	}
	return steamMatchmaking(v)
}

func (s steamMatchmaking) CreateLobby(eLobbyType ELobbyType, cMaxMembers int32) (msg LobbyCreated_t, err error) {
	v, err := theDLL.call(flatAPI_ISteamMatchmaking_CreateLobby, uintptr(s), uintptr(eLobbyType), uintptr(cMaxMembers))
	if err != nil {
		panic(err)
	}
	apiHandle := SteamAPICallbackHandle(v)
	data := getAPICallResult(apiHandle)
	if data != nil {
		created := (*LobbyCreated_t)(unsafe.Pointer(&data[0]))
		return *created, nil
	} else {
		return msg, errors.New("could not complete CreateLobby call")
	}
}

func (s steamMatchmaking) RequestLobbyList() (list LobbyMatchList_t, err error) {
	v, err := theDLL.call(flatAPI_ISteamMatchmaking_RequestLobbyList, uintptr(s))
	if err != nil {
		panic(err)
	}
	apiHandle := SteamAPICallbackHandle(v)
	data := getAPICallResult(apiHandle)
	if data != nil {
		created := (*LobbyMatchList_t)(unsafe.Pointer(&data[0]))
		return *created, nil
	} else {
		return list, errors.New("could not complete RequestLobbyList call")
	}
}

func (s steamMatchmaking) LeaveLobby(steamIDLobby CSteamID) {
	_, err := theDLL.call(flatAPI_ISteamMatchmaking_LeaveLobby, uintptr(s), uintptr(steamIDLobby))
	if err != nil {
		panic(err)
	}
}

func (s steamMatchmaking) GetLobbyByIndex(iLobby int32) CSteamID {
	v, err := theDLL.call(flatAPI_ISteamMatchmaking_GetLobbyByIndex, uintptr(s), uintptr(iLobby))
	if err != nil {
		panic(err)
	}
	return CSteamID(v)
}

func getHSteamPipe() HSteamPipe {
	v, err := theDLL.call(SteamAPI_GetHSteamPipe)
	if err != nil {
		panic(err)
	}
	return HSteamPipe(v)
}

func manualDispatch_Init() {
	_, err := theDLL.call(SteamAPI_ManualDispatch_Init)
	if err != nil {
		panic(err)
	}
}

func manualDispatch_RunFrame(hSteamPipe HSteamPipe) {
	_, err := theDLL.call(SteamAPI_ManualDispatch_RunFrame)
	if err != nil {
		panic(err)
	}

}

func manualDispatch_GetNextCallback(hSteamPipe HSteamPipe) (CallbackMsg_t, bool) {
	var callback CallbackMsg_t
	v, err := theDLL.call(SteamAPI_ManualDispatch_GetNextCallback, uintptr(hSteamPipe), uintptr(unsafe.Pointer(&callback)))
	if err != nil {
		panic(err)
	}
	return callback, v != 0
}

func manualDispatch_FreeLastCallback(hSteamPipe HSteamPipe) {
	_, err := theDLL.call(SteamAPI_ManualDispatch_FreeLastCallback)
	if err != nil {
		panic(err)
	}

}

func manualDispatch_GetAPICallResult(hSteamPipe HSteamPipe, hSteamAPICall SteamAPICallbackHandle,
	cubCallback int32, iCallbackExpected int32) ([]byte, bool) {
	var failed int
	data := make([]byte, cubCallback)

	v, err := theDLL.call(SteamAPI_ManualDispatch_GetAPICallResult,
		uintptr(hSteamPipe), uintptr(hSteamAPICall),
		uintptr(unsafe.Pointer(&data[0])), uintptr(cubCallback), uintptr(iCallbackExpected), uintptr(failed))
	if err != nil {
		panic(err)
	}
	return data, byte(v) != 0

}

func getAPICallResult(handle SteamAPICallbackHandle) []byte {
	hSteamPipe := getHSteamPipe()
	manualDispatch_RunFrame(hSteamPipe)
	for {
		callback, available := manualDispatch_GetNextCallback(hSteamPipe)
		if !available {
			break
		}

		if callback.m_iCallback == int32(k_iSteamAPICallbackCallCompleted) {
			pCallCompleted := (*SteamAPICallCompleted_t)(unsafe.Pointer(callback.m_pubParam))
			if pCallCompleted.m_hAsyncCall != handle {
				continue
			}

			if data, success := manualDispatch_GetAPICallResult(hSteamPipe, pCallCompleted.m_hAsyncCall, callback.m_cubParam, callback.m_iCallback); success {
				manualDispatch_FreeLastCallback(hSteamPipe)
				return data
			}
		}

		manualDispatch_FreeLastCallback(hSteamPipe)
	}
	return nil
}
