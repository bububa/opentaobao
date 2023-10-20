package cntms

import (
	"sync"
)

// CnTmsLogisticsOrderGotService 结构体
type CnTmsLogisticsOrderGotService struct {
	// 揽收日期yyyyMMdd
	GotDate string `json:"got_date,omitempty" xml:"got_date,omitempty"`
	// 揽收时间段 09:00-10:00
	GotRange string `json:"got_range,omitempty" xml:"got_range,omitempty"`
}

var poolCnTmsLogisticsOrderGotService = sync.Pool{
	New: func() any {
		return new(CnTmsLogisticsOrderGotService)
	},
}

// GetCnTmsLogisticsOrderGotService() 从对象池中获取CnTmsLogisticsOrderGotService
func GetCnTmsLogisticsOrderGotService() *CnTmsLogisticsOrderGotService {
	return poolCnTmsLogisticsOrderGotService.Get().(*CnTmsLogisticsOrderGotService)
}

// ReleaseCnTmsLogisticsOrderGotService 释放CnTmsLogisticsOrderGotService
func ReleaseCnTmsLogisticsOrderGotService(v *CnTmsLogisticsOrderGotService) {
	v.GotDate = ""
	v.GotRange = ""
	poolCnTmsLogisticsOrderGotService.Put(v)
}
