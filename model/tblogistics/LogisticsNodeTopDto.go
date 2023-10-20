package tblogistics

import (
	"sync"
)

// LogisticsNodeTopDto 结构体
type LogisticsNodeTopDto struct {
	// ACCEPT(已揽收),TRANSPORT(运输中),DELIVERING(派送中),SIGN(已签收),CANCEL(已取消),FAILED(物流异常)
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 操作时间戳，精确到毫秒（ms）
	OperateTime int64 `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 配送员信息
	Delivery *DeliveryTopDto `json:"delivery,omitempty" xml:"delivery,omitempty"`
	// 货物所在的当前位置
	Location *LocationTopDto `json:"location,omitempty" xml:"location,omitempty"`
}

var poolLogisticsNodeTopDto = sync.Pool{
	New: func() any {
		return new(LogisticsNodeTopDto)
	},
}

// GetLogisticsNodeTopDto() 从对象池中获取LogisticsNodeTopDto
func GetLogisticsNodeTopDto() *LogisticsNodeTopDto {
	return poolLogisticsNodeTopDto.Get().(*LogisticsNodeTopDto)
}

// ReleaseLogisticsNodeTopDto 释放LogisticsNodeTopDto
func ReleaseLogisticsNodeTopDto(v *LogisticsNodeTopDto) {
	v.Action = ""
	v.OperateTime = 0
	v.Delivery = nil
	v.Location = nil
	poolLogisticsNodeTopDto.Put(v)
}
