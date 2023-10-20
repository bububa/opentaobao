package tbk

import (
	"sync"
)

// ServiceFeeDto 结构体
type ServiceFeeDto struct {
	// 专项服务费率（字段已废弃）
	ShareRelativeRate string `json:"share_relative_rate,omitempty" xml:"share_relative_rate,omitempty"`
	// 结算专项服务费（字段已废弃）
	ShareFee string `json:"share_fee,omitempty" xml:"share_fee,omitempty"`
	// 预估专项服务费（字段已废弃）
	SharePreFee string `json:"share_pre_fee,omitempty" xml:"share_pre_fee,omitempty"`
	// 专项服务费来源，122-渠道（字段已废弃）
	TkShareRoleType int64 `json:"tk_share_role_type,omitempty" xml:"tk_share_role_type,omitempty"`
}

var poolServiceFeeDto = sync.Pool{
	New: func() any {
		return new(ServiceFeeDto)
	},
}

// GetServiceFeeDto() 从对象池中获取ServiceFeeDto
func GetServiceFeeDto() *ServiceFeeDto {
	return poolServiceFeeDto.Get().(*ServiceFeeDto)
}

// ReleaseServiceFeeDto 释放ServiceFeeDto
func ReleaseServiceFeeDto(v *ServiceFeeDto) {
	v.ShareRelativeRate = ""
	v.ShareFee = ""
	v.SharePreFee = ""
	v.TkShareRoleType = 0
	poolServiceFeeDto.Put(v)
}
