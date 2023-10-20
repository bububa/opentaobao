package hotel

import (
	"sync"
)

// HotelInfoParam 结构体
type HotelInfoParam struct {
	// pid
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 经过混淆的userId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 上次请求的分页数据中最后一个shid
	LastShid int64 `json:"last_shid,omitempty" xml:"last_shid,omitempty"`
	// 每页数据量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 单个请求的shid
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolHotelInfoParam = sync.Pool{
	New: func() any {
		return new(HotelInfoParam)
	},
}

// GetHotelInfoParam() 从对象池中获取HotelInfoParam
func GetHotelInfoParam() *HotelInfoParam {
	return poolHotelInfoParam.Get().(*HotelInfoParam)
}

// ReleaseHotelInfoParam 释放HotelInfoParam
func ReleaseHotelInfoParam(v *HotelInfoParam) {
	v.Pid = ""
	v.OpenId = ""
	v.CityCode = 0
	v.LastShid = 0
	v.PageSize = 0
	v.Shid = 0
	poolHotelInfoParam.Put(v)
}
