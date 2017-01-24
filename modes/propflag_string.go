// Code generated by "stringer -type=PropFlag modes.go"; DO NOT EDIT

package modes

import "fmt"

const _PropFlag_name = "PropNonePropMutePropPrivatePropInvitePropTopicrestPropInternalPropSecretPropNoCTCPPropNoActionPropNoKicksPropNoCapsPropNoRejoinPropLargelistPropNoOperKickPropOperonlyPropPermanentPropDisforwardPropNoNoticePropNoColorPropNoNicksPropFreeinvitePropHidebansPropOpmodPropFreefwdPropNoRepeat"

var _PropFlag_map = map[PropFlag]string{
	1:        _PropFlag_name[0:8],
	2:        _PropFlag_name[8:16],
	4:        _PropFlag_name[16:27],
	8:        _PropFlag_name[27:37],
	16:       _PropFlag_name[37:50],
	32:       _PropFlag_name[50:62],
	64:       _PropFlag_name[62:72],
	128:      _PropFlag_name[72:82],
	256:      _PropFlag_name[82:94],
	512:      _PropFlag_name[94:105],
	1024:     _PropFlag_name[105:115],
	2048:     _PropFlag_name[115:127],
	4096:     _PropFlag_name[127:140],
	8192:     _PropFlag_name[140:154],
	16384:    _PropFlag_name[154:166],
	32768:    _PropFlag_name[166:179],
	65536:    _PropFlag_name[179:193],
	131072:   _PropFlag_name[193:205],
	262144:   _PropFlag_name[205:216],
	524288:   _PropFlag_name[216:227],
	1048576:  _PropFlag_name[227:241],
	2097152:  _PropFlag_name[241:253],
	4194304:  _PropFlag_name[253:262],
	8388608:  _PropFlag_name[262:273],
	16777216: _PropFlag_name[273:285],
}

func (i PropFlag) String() string {
	if str, ok := _PropFlag_map[i]; ok {
		return str
	}
	return fmt.Sprintf("PropFlag(%d)", i)
}