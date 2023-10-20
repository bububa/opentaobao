package btrip

import (
	"sync"
)

// TransferInfo 结构体
type TransferInfo struct {
	// 第二程退改签规则列表
	TransferFlightRuleList []FlightRule `json:"transfer_flight_rule_list,omitempty" xml:"transfer_flight_rule_list>flight_rule,omitempty"`
	// 航班号，二程
	TransferFlightNo string `json:"transfer_flight_no,omitempty" xml:"transfer_flight_no,omitempty"`
	// 中转到达时间 (第一程到达时间)
	TransferArrDate string `json:"transfer_arr_date,omitempty" xml:"transfer_arr_date,omitempty"`
	// 中转起飞时间 (第二程起飞时间)
	TransferDepDate string `json:"transfer_dep_date,omitempty" xml:"transfer_dep_date,omitempty"`
	// 机型
	FlightSize string `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	// 机型代码
	FlightType string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	//  中转到达机场信息
	TransferArrAirportInfo *AirportInfo `json:"transfer_arr_airport_info,omitempty" xml:"transfer_arr_airport_info,omitempty"`
	// 中转出发机场信息
	TransferDepAirportInfo *AirportInfo `json:"transfer_dep_airport_info,omitempty" xml:"transfer_dep_airport_info,omitempty"`
	// 第二程销售航司信息
	TransferAirlineInfo *AirlineInfo `json:"transfer_airline_info,omitempty" xml:"transfer_airline_info,omitempty"`
}

var poolTransferInfo = sync.Pool{
	New: func() any {
		return new(TransferInfo)
	},
}

// GetTransferInfo() 从对象池中获取TransferInfo
func GetTransferInfo() *TransferInfo {
	return poolTransferInfo.Get().(*TransferInfo)
}

// ReleaseTransferInfo 释放TransferInfo
func ReleaseTransferInfo(v *TransferInfo) {
	v.TransferFlightRuleList = v.TransferFlightRuleList[:0]
	v.TransferFlightNo = ""
	v.TransferArrDate = ""
	v.TransferDepDate = ""
	v.FlightSize = ""
	v.FlightType = ""
	v.TransferArrAirportInfo = nil
	v.TransferDepAirportInfo = nil
	v.TransferAirlineInfo = nil
	poolTransferInfo.Put(v)
}
