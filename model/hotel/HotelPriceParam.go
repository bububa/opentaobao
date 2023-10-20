package hotel

import (
	"sync"
)

// HotelPriceParam 结构体
type HotelPriceParam struct {
	// 版本控制(3.0支持信用住)
	DataVersion string `json:"data_version,omitempty" xml:"data_version,omitempty"`
	// 离店日期
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 经过混淆的userId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 标志参数pid
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 标准酒店id和城市编码拼接字符串，最多50个
	ShidCityCode string `json:"shid_city_code,omitempty" xml:"shid_city_code,omitempty"`
	// 入住日期
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 请求类型：0批量请求；1实时请求
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolHotelPriceParam = sync.Pool{
	New: func() any {
		return new(HotelPriceParam)
	},
}

// GetHotelPriceParam() 从对象池中获取HotelPriceParam
func GetHotelPriceParam() *HotelPriceParam {
	return poolHotelPriceParam.Get().(*HotelPriceParam)
}

// ReleaseHotelPriceParam 释放HotelPriceParam
func ReleaseHotelPriceParam(v *HotelPriceParam) {
	v.DataVersion = ""
	v.EndDate = ""
	v.OpenId = ""
	v.Pid = ""
	v.ShidCityCode = ""
	v.StartDate = ""
	v.Type = 0
	poolHotelPriceParam.Put(v)
}
