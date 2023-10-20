package xhotelitem

import (
	"sync"
)

// Coordinate 结构体
type Coordinate struct {
	// 飞猪城市中文名称
	CityCnName string `json:"city_cn_name,omitempty" xml:"city_cn_name,omitempty"`
	// 飞猪城市英文名称
	CityEnName string `json:"city_en_name,omitempty" xml:"city_en_name,omitempty"`
	// 外部经纬度标识id，可以是酒店或城市的id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 批次号
	BatchId int64 `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 飞猪城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 飞猪国家编码
	Country int64 `json:"country,omitempty" xml:"country,omitempty"`
}

var poolCoordinate = sync.Pool{
	New: func() any {
		return new(Coordinate)
	},
}

// GetCoordinate() 从对象池中获取Coordinate
func GetCoordinate() *Coordinate {
	return poolCoordinate.Get().(*Coordinate)
}

// ReleaseCoordinate 释放Coordinate
func ReleaseCoordinate(v *Coordinate) {
	v.CityCnName = ""
	v.CityEnName = ""
	v.OuterId = ""
	v.Latitude = ""
	v.Longitude = ""
	v.BatchId = 0
	v.City = 0
	v.Country = 0
	poolCoordinate.Put(v)
}
