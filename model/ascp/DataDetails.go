package ascp

import (
	"sync"
)

// DataDetails 结构体
type DataDetails struct {
	// 只会返回错误的地址id，如果淘天物流无法识别的地址则返回错误码
	AddressIdResults []AddressIdResults `json:"address_id_results,omitempty" xml:"address_id_results>address_id_results,omitempty"`
	// 只会返回错误的地址信息，如果淘天物流无法识别的地址则返回错误码
	AddressNameResults []AddressNameResults `json:"address_name_results,omitempty" xml:"address_name_results>address_name_results,omitempty"`
	// 只会返回错误的地理围栏信息
	RegionIdResults []RegionIdResults `json:"region_id_results,omitempty" xml:"region_id_results>region_id_results,omitempty"`
	// 记录行ID
	OrderProcessReportId string `json:"order_process_report_id,omitempty" xml:"order_process_report_id,omitempty"`
	// 记录行详细描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 记录行同步结果   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDataDetails = sync.Pool{
	New: func() any {
		return new(DataDetails)
	},
}

// GetDataDetails() 从对象池中获取DataDetails
func GetDataDetails() *DataDetails {
	return poolDataDetails.Get().(*DataDetails)
}

// ReleaseDataDetails 释放DataDetails
func ReleaseDataDetails(v *DataDetails) {
	v.AddressIdResults = v.AddressIdResults[:0]
	v.AddressNameResults = v.AddressNameResults[:0]
	v.RegionIdResults = v.RegionIdResults[:0]
	v.OrderProcessReportId = ""
	v.Message = ""
	v.Success = false
	poolDataDetails.Put(v)
}
