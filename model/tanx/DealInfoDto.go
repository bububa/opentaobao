package tanx

import (
	"sync"
)

// DealInfoDto 结构体
type DealInfoDto struct {
	// 11
	Addresses []DicDto `json:"addresses,omitempty" xml:"addresses>dic_dto,omitempty"`
	// 100
	SellerSiteNames []string `json:"seller_site_names,omitempty" xml:"seller_site_names>string,omitempty"`
	// 100
	DspIds []int64 `json:"dsp_ids,omitempty" xml:"dsp_ids>int64,omitempty"`
	// 100
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty" xml:"advertiser_ids>int64,omitempty"`
	// 100
	IntervalIds []string `json:"interval_ids,omitempty" xml:"interval_ids>string,omitempty"`
	// 100
	Pids []string `json:"pids,omitempty" xml:"pids>string,omitempty"`
	// 100
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 100
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 100
	DealId int64 `json:"deal_id,omitempty" xml:"deal_id,omitempty"`
	// 100
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 100
	DealType int64 `json:"deal_type,omitempty" xml:"deal_type,omitempty"`
	// 100
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolDealInfoDto = sync.Pool{
	New: func() any {
		return new(DealInfoDto)
	},
}

// GetDealInfoDto() 从对象池中获取DealInfoDto
func GetDealInfoDto() *DealInfoDto {
	return poolDealInfoDto.Get().(*DealInfoDto)
}

// ReleaseDealInfoDto 释放DealInfoDto
func ReleaseDealInfoDto(v *DealInfoDto) {
	v.Addresses = v.Addresses[:0]
	v.SellerSiteNames = v.SellerSiteNames[:0]
	v.DspIds = v.DspIds[:0]
	v.AdvertiserIds = v.AdvertiserIds[:0]
	v.IntervalIds = v.IntervalIds[:0]
	v.Pids = v.Pids[:0]
	v.BeginTime = ""
	v.EndTime = ""
	v.DealId = 0
	v.Price = 0
	v.DealType = 0
	v.Status = 0
	poolDealInfoDto.Put(v)
}
