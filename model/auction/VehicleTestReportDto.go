package auction

import (
	"sync"
)

// VehicleTestReportDto 结构体
type VehicleTestReportDto struct {
	// 复核vin码
	ReviewVin string `json:"review_vin,omitempty" xml:"review_vin,omitempty"`
	// 原始vin码（拍卖传的vin码，没传返回空）
	OriginVin string `json:"origin_vin,omitempty" xml:"origin_vin,omitempty"`
	// 机动车报告地址链接
	ReportUrl string `json:"report_url,omitempty" xml:"report_url,omitempty"`
	// 机动车机构化属性，json格式
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 时间戳
	Datestamp int64 `json:"datestamp,omitempty" xml:"datestamp,omitempty"`
	// 是否有生成报告，true 有，false 未生成
	HasReport bool `json:"has_report,omitempty" xml:"has_report,omitempty"`
}

var poolVehicleTestReportDto = sync.Pool{
	New: func() any {
		return new(VehicleTestReportDto)
	},
}

// GetVehicleTestReportDto() 从对象池中获取VehicleTestReportDto
func GetVehicleTestReportDto() *VehicleTestReportDto {
	return poolVehicleTestReportDto.Get().(*VehicleTestReportDto)
}

// ReleaseVehicleTestReportDto 释放VehicleTestReportDto
func ReleaseVehicleTestReportDto(v *VehicleTestReportDto) {
	v.ReviewVin = ""
	v.OriginVin = ""
	v.ReportUrl = ""
	v.Attribute = ""
	v.Datestamp = 0
	v.HasReport = false
	poolVehicleTestReportDto.Put(v)
}
