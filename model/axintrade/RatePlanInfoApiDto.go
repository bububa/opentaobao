package axintrade

import (
	"sync"
)

// RatePlanInfoApiDto 结构体
type RatePlanInfoApiDto struct {
	// 每间房rate信息
	RateUnitList []RateUnitDto `json:"rate_unit_list,omitempty" xml:"rate_unit_list>rate_unit_dto,omitempty"`
	// 最晚到店时间
	LatestCheckOutTime string `json:"latest_check_out_time,omitempty" xml:"latest_check_out_time,omitempty"`
	// 总房价
	TotalRoomPrice string `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 床型描述
	BedTypeDesc string `json:"bed_type_desc,omitempty" xml:"bed_type_desc,omitempty"`
	// 最早可以办理入住时间
	EarliestCheckInTime string `json:"earliest_check_in_time,omitempty" xml:"earliest_check_in_time,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 最大订购数量
	MaxBookingNum int64 `json:"max_booking_num,omitempty" xml:"max_booking_num,omitempty"`
	// 每间房最大可入住人数
	MaxOccupancyNum int64 `json:"max_occupancy_num,omitempty" xml:"max_occupancy_num,omitempty"`
	// 取消政策
	CancelPolicy *CancelPolicyDto `json:"cancel_policy,omitempty" xml:"cancel_policy,omitempty"`
	// 最大库存量
	MaxInventory int64 `json:"max_inventory,omitempty" xml:"max_inventory,omitempty"`
	// 人民币总金额
	CnyTotalPrice int64 `json:"cny_total_price,omitempty" xml:"cny_total_price,omitempty"`
	// 汇率
	ExchangeRate float64 `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 0-全日房, 1-小时房
	RpType int64 `json:"rp_type,omitempty" xml:"rp_type,omitempty"`
	// 小时房到店时间&amp;连住时长
	ArrivalTime *ArrivalTimeDto `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// 是否即时确认
	InstantConfirm bool `json:"instant_confirm,omitempty" xml:"instant_confirm,omitempty"`
}

var poolRatePlanInfoApiDto = sync.Pool{
	New: func() any {
		return new(RatePlanInfoApiDto)
	},
}

// GetRatePlanInfoApiDto() 从对象池中获取RatePlanInfoApiDto
func GetRatePlanInfoApiDto() *RatePlanInfoApiDto {
	return poolRatePlanInfoApiDto.Get().(*RatePlanInfoApiDto)
}

// ReleaseRatePlanInfoApiDto 释放RatePlanInfoApiDto
func ReleaseRatePlanInfoApiDto(v *RatePlanInfoApiDto) {
	v.RateUnitList = v.RateUnitList[:0]
	v.LatestCheckOutTime = ""
	v.TotalRoomPrice = ""
	v.BedTypeDesc = ""
	v.EarliestCheckInTime = ""
	v.CurrencyCode = ""
	v.MaxBookingNum = 0
	v.MaxOccupancyNum = 0
	v.CancelPolicy = nil
	v.MaxInventory = 0
	v.CnyTotalPrice = 0
	v.ExchangeRate = 0
	v.RpType = 0
	v.ArrivalTime = nil
	v.InstantConfirm = false
	poolRatePlanInfoApiDto.Put(v)
}
