// generated by stringer -type=UserFlag; DO NOT EDIT

package modes

import "fmt"

const _UserFlag_name = "UPropNoneUPropInvisibleUPropCalleridUPropIrcopUPropCloakedUPropAdminUPropOverrideUPropNoCTCPUPropDeafUPropDisforwardUPropRegpmUPropSoftcallUPropNoinviteUPropNostalkUPropSSLClientUPropNetworkService"

var _UserFlag_index = [...]uint8{0, 9, 23, 36, 46, 58, 68, 81, 92, 101, 116, 126, 139, 152, 164, 178, 197}

func (i UserFlag) String() string {
	if i < 0 || i+1 >= UserFlag(len(_UserFlag_index)) {
		return fmt.Sprintf("UserFlag(%d)", i)
	}
	return _UserFlag_name[_UserFlag_index[i]:_UserFlag_index[i+1]]
}
