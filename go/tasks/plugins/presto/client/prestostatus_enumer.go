// Code generated by "enumer --type=PrestoStatus"; DO NOT EDIT.

//
package client

import (
	"fmt"
)

const _PrestoStatusName = "PrestoStatusUnknownPrestoStatusWaitingPrestoStatusRunningPrestoStatusFinishedPrestoStatusFailedPrestoStatusCancelled"

var _PrestoStatusIndex = [...]uint8{0, 19, 38, 57, 77, 95, 116}

func (i PrestoStatus) String() string {
	if i < 0 || i >= PrestoStatus(len(_PrestoStatusIndex)-1) {
		return fmt.Sprintf("PrestoStatus(%d)", i)
	}
	return _PrestoStatusName[_PrestoStatusIndex[i]:_PrestoStatusIndex[i+1]]
}

var _PrestoStatusValues = []PrestoStatus{0, 1, 2, 3, 4, 5}

var _PrestoStatusNameToValueMap = map[string]PrestoStatus{
	_PrestoStatusName[0:19]:   0,
	_PrestoStatusName[19:38]:  1,
	_PrestoStatusName[38:57]:  2,
	_PrestoStatusName[57:77]:  3,
	_PrestoStatusName[77:95]:  4,
	_PrestoStatusName[95:116]: 5,
}

// PrestoStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrestoStatusString(s string) (PrestoStatus, error) {
	if val, ok := _PrestoStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PrestoStatus values", s)
}

// PrestoStatusValues returns all values of the enum
func PrestoStatusValues() []PrestoStatus {
	return _PrestoStatusValues
}

// IsAPrestoStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PrestoStatus) IsAPrestoStatus() bool {
	for _, v := range _PrestoStatusValues {
		if i == v {
			return true
		}
	}
	return false
}