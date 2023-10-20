package cloudgame

import (
	"sync"
)

// AlibabaCgameAvatarUserbodyQueryData 结构体
type AlibabaCgameAvatarUserbodyQueryData struct {
	// body traceID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 脸部数据
	FaceDataJson string `json:"face_data_json,omitempty" xml:"face_data_json,omitempty"`
	// 请求唯一标识ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户mixUserId
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// 扩展数据
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 性别， 1-male, 2-female
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
}

var poolAlibabaCgameAvatarUserbodyQueryData = sync.Pool{
	New: func() any {
		return new(AlibabaCgameAvatarUserbodyQueryData)
	},
}

// GetAlibabaCgameAvatarUserbodyQueryData() 从对象池中获取AlibabaCgameAvatarUserbodyQueryData
func GetAlibabaCgameAvatarUserbodyQueryData() *AlibabaCgameAvatarUserbodyQueryData {
	return poolAlibabaCgameAvatarUserbodyQueryData.Get().(*AlibabaCgameAvatarUserbodyQueryData)
}

// ReleaseAlibabaCgameAvatarUserbodyQueryData 释放AlibabaCgameAvatarUserbodyQueryData
func ReleaseAlibabaCgameAvatarUserbodyQueryData(v *AlibabaCgameAvatarUserbodyQueryData) {
	v.TraceId = ""
	v.FaceDataJson = ""
	v.RequestId = ""
	v.MixUserId = ""
	v.ExtInfo = ""
	v.Gender = 0
	poolAlibabaCgameAvatarUserbodyQueryData.Put(v)
}
