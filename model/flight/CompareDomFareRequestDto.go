package flight

import (
	"sync"
)

// CompareDomFareRequestDto 结构体
type CompareDomFareRequestDto struct {
	// 航空公司
	AirlineCodes []string `json:"airline_codes,omitempty" xml:"airline_codes>string,omitempty"`
	// 产品类型：0，普通；1，小团；2，学生；3，青年；4，老年；5，地区；6，会员；10，学生认证；11，年龄；为空表示不限制
	ProductTypes []string `json:"product_types,omitempty" xml:"product_types>string,omitempty"`
	// 航程信息
	OdInfos []OdInfoQueryDto `json:"od_infos,omitempty" xml:"od_infos>od_info_query_dto,omitempty"`
	// 销售方式：0，无；1，打包销售套餐1；2，打包销售套餐2；3，打包销售套餐3；4，返现-航司运价；5，返现-销售方包装；为空表示不限制
	SaleModeCodes []string `json:"sale_mode_codes,omitempty" xml:"sale_mode_codes>string,omitempty"`
	// 航司+航班号，多个航班号用#分隔，航班号范围用-连接，例如：HU9990-HU9995#HU9999
	FlightNoStr string `json:"flight_no_str,omitempty" xml:"flight_no_str,omitempty"`
	// 舱位代码，仅可传入一个值
	CabinCodeStr string `json:"cabin_code_str,omitempty" xml:"cabin_code_str,omitempty"`
	// 行程类型：0，单程；1，往返；暂时仅支持单程
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// 共享航班：0，不支持代码共享；1，支持代码共享；2，仅支持代码共享；为空表示不限制
	SupportCodeShare int64 `json:"support_code_share,omitempty" xml:"support_code_share,omitempty"`
	// 政策投放情况：0，未投放；1，已投放，非最优惠；2，已投放，且为最优惠政策；为空表示不限制
	PolicyDeployStatus int64 `json:"policy_deploy_status,omitempty" xml:"policy_deploy_status,omitempty"`
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}

var poolCompareDomFareRequestDto = sync.Pool{
	New: func() any {
		return new(CompareDomFareRequestDto)
	},
}

// GetCompareDomFareRequestDto() 从对象池中获取CompareDomFareRequestDto
func GetCompareDomFareRequestDto() *CompareDomFareRequestDto {
	return poolCompareDomFareRequestDto.Get().(*CompareDomFareRequestDto)
}

// ReleaseCompareDomFareRequestDto 释放CompareDomFareRequestDto
func ReleaseCompareDomFareRequestDto(v *CompareDomFareRequestDto) {
	v.AirlineCodes = v.AirlineCodes[:0]
	v.ProductTypes = v.ProductTypes[:0]
	v.OdInfos = v.OdInfos[:0]
	v.SaleModeCodes = v.SaleModeCodes[:0]
	v.FlightNoStr = ""
	v.CabinCodeStr = ""
	v.TripType = 0
	v.SupportCodeShare = 0
	v.PolicyDeployStatus = 0
	v.AgentId = 0
	poolCompareDomFareRequestDto.Put(v)
}
