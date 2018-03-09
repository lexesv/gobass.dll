// bass project bass.go
package bass

/*
#cgo linux CFLAGS: -I/usr/include -I.
#cgo linux LDFLAGS: -L${SRCDIR} -L/usr/lib -Wl,-rpath=\$ORIGIN -lbass
#include "bass.h"
*/
import "C"

import (
	"fmt"
	"strconv"
	"unsafe"
	"errors"
)

/*
Init
BOOL BASSDEF(BASS_Init)(int device, DWORD freq, DWORD flags, void *win, void *dsguid);
*/
func Init(device int, freq int, flags int) (bool, error) {
	if C.BASS_Init(C.int(device), C.DWORD(freq), C.DWORD(flags), nil, nil) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
Free
BOOL BASS_Free();
 */
func Free() (bool, error) {
	if C.BASS_Free() != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
GetConfig
DWORD BASSDEF(BASS_GetConfig)(DWORD option);
*/
func GetConfig(option int) (int, error) {
	v := (int)(C.BASS_GetConfig(C.DWORD(option)))
	if v != -1 {
		return v, nil
	} else {
		return -1, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
SetConfig
BOOL BASSDEF(BASS_SetConfig)(DWORD option, DWORD value);
*/
func SetConfig(option, value int) (bool, error) {
	if C.BASS_SetConfig(C.DWORD(option), C.DWORD(value)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// GetVol
// float BASSDEF(BASS_GetVolume)();
func GetVol() (float32, error) {
	if v := C.BASS_GetVolume(); v != -1 {
		return float32(v) * float32(100), nil
	} else {
		return -1, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// SetVol
func SetVol(v float32) (bool, error) {
	if C.BASS_SetVolume(C.float(v/100)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// StreamCreateURL
func StreamCreateURL(url string) (int, error) {
	ch := C.BASS_StreamCreateURL(C.CString(url), 0, C.BASS_STREAM_BLOCK|C.BASS_STREAM_STATUS|C.BASS_STREAM_AUTOFREE, nil, nil)
	if ch != 0 {
		return int(ch), nil
	} else {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// BASS_StreamCreateFile
// HSTREAM BASSDEF(BASS_StreamCreateFile)(BOOL mem, const void *file, QWORD offset, QWORD length, DWORD flags);
func StreamCreateFile(file string) (int, error) {
	ch := C.BASS_StreamCreateFile(0, unsafe.Pointer(C.CString(file)), 0, 0, C.BASS_ASYNCFILE|C.BASS_STREAM_AUTOFREE)
	if ch != 0 {
		return int(ch), nil
	} else {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelPlay
// BOOL BASSDEF(BASS_ChannelPlay)(DWORD handle, BOOL restart);
func ChannelPlay(ch int) (bool, error) {
	if C.BASS_ChannelPlay(C.DWORD(ch), 1) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelPause
// BOOL BASSDEF(BASS_ChannelPause)(DWORD handle);
func ChannelPause(ch int) (bool, error) {
	if C.BASS_ChannelPause(C.DWORD(ch)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelStop
// BOOL BASSDEF(BASS_ChannelStop)(DWORD handle);
func ChannelStop(ch int) (bool, error) {
	if C.BASS_ChannelStop(C.DWORD(ch)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelStatus
// DWORD BASSDEF(BASS_ChannelIsActive)(DWORD handle);
func ChannelStatus(ch int) (c uint, s string) {
	c = uint(C.BASS_ChannelIsActive(C.DWORD(ch)))
	switch c {
	case 0:
		s = "BASS_ACTIVE_STOPPED"
		break
	case 1:
		s = "BASS_ACTIVE_PLAYING"
		break
	case 2:
		s = "BASS_ACTIVE_STALLED"
		break
	case 3:
		s = "BASS_ACTIVE_PAUSED"
		break
	}
	return c, s
}

// ChannelGetAttribute
// BOOL BASSDEF(BASS_ChannelGetAttribute)(DWORD handle, DWORD attrib, float *value);
func ChannelGetAttribute(ch int, attrib int) (float32, error) {
	var cvalue *C.float
	if C.BASS_ChannelGetAttribute(C.DWORD(ch), C.DWORD(attrib), cvalue) != 0 {
		if v, err := strconv.ParseFloat(fmt.Sprintf("%v", cvalue), 32); err != nil {
			return -1, nil
		} else {
			return float32(v), nil
		}
	} else {
		return -1, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelSetAttribute
// BOOL BASSDEF(BASS_ChannelSetAttribute)(DWORD handle, DWORD attrib, float value);
func ChannelSetAttribute(ch int, attrib int, value float32) (bool, error) {
	if C.BASS_ChannelSetAttribute(C.DWORD(ch), C.DWORD(attrib), C.float(value)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelGetVolume
// value: 0-100
func ChannelGetVolume(ch int) (float32, error) {
	v, err := ChannelGetAttribute(ch, BASS_ATTRIB_VOL)
	if v > 0 {
		v = v * 100
	}
	return v, err
}

// ChannelSetVolume
// value: 0-100
func ChannelSetVolume(ch int, value float32) (bool, error) {
	return ChannelSetAttribute(ch, BASS_ATTRIB_VOL, value/100)
}

// ChannelGetTags
// const char *BASSDEF(BASS_ChannelGetTags)(DWORD handle, DWORD tags);
func ChannelGetTags(ch int, tag int) string {
	return C.GoString(C.BASS_ChannelGetTags(C.DWORD(ch), C.DWORD(tag)))
}

// PluginLoad
func PluginLoad(file string) (handle int, err error) {
	if h := C.BASS_PluginLoad(C.CString(file), 0); h != 0 {
		return int(h), nil
	} else {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// PluginFree
func PluginFree(handle int) (bool, error) {
	if C.BASS_PluginFree(C.HPLUGIN(handle)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/* RECORD */
/*
RecordInit
BOOL BASSDEF(BASS_RecordInit)(int device);
*/
func RecordInit(device int) (bool, error) {
	if C.BASS_RecordInit(C.int(device)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
RecordFree
BOOL BASSDEF(BASS_RecordFree)();
 */
func RecordFree() (bool, error) {
	if C.BASS_RecordFree() != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
RecordStart
HRECORD BASSDEF(BASS_RecordStart)(DWORD freq, DWORD chans, DWORD flags, RECORDPROC *proc, void *user);
 */

func RecordStart(freq int, chans int, flags int, proc RecordCallback) (int, error) {
	h := C.BASS_RecordStart(C.DWORD(freq), C.DWORD(chans), C.DWORD(flags), (*C.RECORDPROC)(unsafe.Pointer(&proc)), nil)
	//h := C.BASS_RecordStart(C.DWORD(freq), C.DWORD(chans), C.DWORD(flags), nil, nil)
	if h != 0 {
		return int(h), nil
	} else {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

//typedef BOOL (CALLBACK RECORDPROC)(HRECORD handle, const void *buffer, DWORD length, void *user);
type RecordCallback = func(handle C.HRECORD, buffer *C.char, length C.DWORD, user *C.char) bool

func errMsg(c int) error {
	codes := make(map[int]string)
	codes[0] = "all is OK"
	codes[1] = "memory error"
	codes[2] = "can't open the file"
	codes[3] = "can't find a free/valid driver"
	codes[4] = "the sample buffer was lost"
	codes[5] = "invalid handle"
	codes[6] = "unsupported sample format"
	codes[7] = "invalid position"
	codes[8] = "BASS_Init has not been successfully called"
	codes[9] = "BASS_Start has not been successfully called"
	codes[10] = "SSL/HTTPS support isn't available"
	codes[14] = "already initialized/paused/whatever"
	codes[18] = "can't get a free channel"
	codes[19] = "an illegal type was specified"
	codes[20] = "an illegal parameter was specified"
	codes[21] = "no 3D support"
	codes[22] = "no EAX support"
	codes[23] = "illegal device number"
	codes[24] = "not playing"
	codes[25] = "illegal sample rate"
	codes[27] = "the stream is not a file stream"
	codes[29] = "no hardware voices available"
	codes[31] = "the MOD music has no sequence data"
	codes[32] = "no internet connection could be opened"
	codes[33] = "couldn't create the file"
	codes[34] = "effects are not available"
	codes[37] = "requested data is not available"
	codes[38] = "the channel is/isn't a 'decoding channel'"
	codes[39] = "a sufficient DirectX version is not installed"
	codes[40] = "connection timedout"
	codes[41] = "unsupported file format"
	codes[42] = "unavailable speaker"
	codes[43] = "invalid BASS version (used by add-ons)"
	codes[44] = "codec is not available/supported"
	codes[45] = "the channel/file has ended"
	codes[46] = "the device is busy"
	codes[-1] = "some other mystery problem"
	return errors.New(codes[c])
}
