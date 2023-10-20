package larkiot

import (
	"sync"
)

// CinemaIotResponse 结构体
type CinemaIotResponse struct {
	// 影院内码
	CinemaLinkId string `json:"cinema_link_id,omitempty" xml:"cinema_link_id,omitempty"`
	// 影院名称
	CinemaName string `json:"cinema_name,omitempty" xml:"cinema_name,omitempty"`
	// 影院国家编码
	CinemaId string `json:"cinema_id,omitempty" xml:"cinema_id,omitempty"`
	// 城市编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
}

var poolCinemaIotResponse = sync.Pool{
	New: func() any {
		return new(CinemaIotResponse)
	},
}

// GetCinemaIotResponse() 从对象池中获取CinemaIotResponse
func GetCinemaIotResponse() *CinemaIotResponse {
	return poolCinemaIotResponse.Get().(*CinemaIotResponse)
}

// ReleaseCinemaIotResponse 释放CinemaIotResponse
func ReleaseCinemaIotResponse(v *CinemaIotResponse) {
	v.CinemaLinkId = ""
	v.CinemaName = ""
	v.CinemaId = ""
	v.CityCode = ""
	v.City = ""
	poolCinemaIotResponse.Put(v)
}
