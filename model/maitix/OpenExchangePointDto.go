package maitix

import (
	"sync"
)

// OpenExchangePointDto 结构体
type OpenExchangePointDto struct {
	// 换票点名称
	PointName string `json:"point_name,omitempty" xml:"point_name,omitempty"`
	// 详细地址
	PointAddr string `json:"point_addr,omitempty" xml:"point_addr,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 营业时间描述
	BizTimeShow string `json:"biz_time_show,omitempty" xml:"biz_time_show,omitempty"`
	// 提示信息备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 类型 1临时换票点 2固定换票点
	PointType string `json:"point_type,omitempty" xml:"point_type,omitempty"`
	// 换票形式 1 pc 2 自助机
	ExchangeType string `json:"exchange_type,omitempty" xml:"exchange_type,omitempty"`
	// 有效期
	ValidPeriod string `json:"valid_period,omitempty" xml:"valid_period,omitempty"`
	// 换票点id
	PointId int64 `json:"point_id,omitempty" xml:"point_id,omitempty"`
}

var poolOpenExchangePointDto = sync.Pool{
	New: func() any {
		return new(OpenExchangePointDto)
	},
}

// GetOpenExchangePointDto() 从对象池中获取OpenExchangePointDto
func GetOpenExchangePointDto() *OpenExchangePointDto {
	return poolOpenExchangePointDto.Get().(*OpenExchangePointDto)
}

// ReleaseOpenExchangePointDto 释放OpenExchangePointDto
func ReleaseOpenExchangePointDto(v *OpenExchangePointDto) {
	v.PointName = ""
	v.PointAddr = ""
	v.Longitude = ""
	v.Latitude = ""
	v.BizTimeShow = ""
	v.Remark = ""
	v.PointType = ""
	v.ExchangeType = ""
	v.ValidPeriod = ""
	v.PointId = 0
	poolOpenExchangePointDto.Put(v)
}
