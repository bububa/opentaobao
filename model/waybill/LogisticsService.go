package waybill

import (
	"sync"
)

// LogisticsService 结构体
type LogisticsService struct {
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务类型值，json格式表示
	ServiceValue4Json string `json:"service_value4_json,omitempty" xml:"service_value4_json,omitempty"`
}

var poolLogisticsService = sync.Pool{
	New: func() any {
		return new(LogisticsService)
	},
}

// GetLogisticsService() 从对象池中获取LogisticsService
func GetLogisticsService() *LogisticsService {
	return poolLogisticsService.Get().(*LogisticsService)
}

// ReleaseLogisticsService 释放LogisticsService
func ReleaseLogisticsService(v *LogisticsService) {
	v.ServiceCode = ""
	v.ServiceValue4Json = ""
	poolLogisticsService.Put(v)
}
