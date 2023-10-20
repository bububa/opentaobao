package btrip

import (
	"sync"
)

// HisvResultSupport 结构体
type HisvResultSupport struct {
	// 出参
	ModuleList []OpenApiZzdFlightOrderRs `json:"module_list,omitempty" xml:"module_list>open_api_zzd_flight_order_rs,omitempty"`
	// 出参
	TradeList []OpenApiZzdHotelOrderRs `json:"trade_list,omitempty" xml:"trade_list>open_api_zzd_hotel_order_rs,omitempty"`
	// 结果msg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// module
	Module *OpenIsvBillSettlementRs `json:"module,omitempty" xml:"module,omitempty"`
	// 结果code
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolHisvResultSupport = sync.Pool{
	New: func() any {
		return new(HisvResultSupport)
	},
}

// GetHisvResultSupport() 从对象池中获取HisvResultSupport
func GetHisvResultSupport() *HisvResultSupport {
	return poolHisvResultSupport.Get().(*HisvResultSupport)
}

// ReleaseHisvResultSupport 释放HisvResultSupport
func ReleaseHisvResultSupport(v *HisvResultSupport) {
	v.ModuleList = v.ModuleList[:0]
	v.TradeList = v.TradeList[:0]
	v.ResultMsg = ""
	v.Module = nil
	v.ResultCode = 0
	v.Success = false
	poolHisvResultSupport.Put(v)
}
