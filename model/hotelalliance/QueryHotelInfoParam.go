package hotelalliance

import (
	"sync"
)

// QueryHotelInfoParam 结构体
type QueryHotelInfoParam struct {
	// 飞猪卖家酒店id
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 单体联盟飞猪卖家酒店id
	Ahid int64 `json:"ahid,omitempty" xml:"ahid,omitempty"`
}

var poolQueryHotelInfoParam = sync.Pool{
	New: func() any {
		return new(QueryHotelInfoParam)
	},
}

// GetQueryHotelInfoParam() 从对象池中获取QueryHotelInfoParam
func GetQueryHotelInfoParam() *QueryHotelInfoParam {
	return poolQueryHotelInfoParam.Get().(*QueryHotelInfoParam)
}

// ReleaseQueryHotelInfoParam 释放QueryHotelInfoParam
func ReleaseQueryHotelInfoParam(v *QueryHotelInfoParam) {
	v.Hid = 0
	v.Ahid = 0
	poolQueryHotelInfoParam.Put(v)
}
