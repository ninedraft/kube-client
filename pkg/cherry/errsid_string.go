// Code generated by "stringer -type ErrSID"; DO NOT EDIT.

package cherry

import "strconv"

const _ErrSID_name = "UninitializedSID"

var _ErrSID_index = [...]uint8{0, 16}

func (i ErrSID) String() string {
	if i >= ErrSID(len(_ErrSID_index)-1) {
		return "ErrSID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrSID_name[_ErrSID_index[i]:_ErrSID_index[i+1]]
}
