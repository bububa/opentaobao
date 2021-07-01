package tanx

// DealInfoDto 结构体
type DealInfoDto struct {
	// 11
	Addresses []DicDto `json:"addresses,omitempty" xml:"addresses>dic_dto,omitempty"`
	// 100
	DealId int64 `json:"deal_id,omitempty" xml:"deal_id,omitempty"`
	// 100
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 100
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 100
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 100
	DealType int64 `json:"deal_type,omitempty" xml:"deal_type,omitempty"`
	// 100
	SellerSiteNames []string `json:"seller_site_names,omitempty" xml:"seller_site_names>string,omitempty"`
	// 100
	DspIds []int64 `json:"dsp_ids,omitempty" xml:"dsp_ids>int64,omitempty"`
	// 100
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty" xml:"advertiser_ids>int64,omitempty"`
	// 100
	IntervalIds []string `json:"interval_ids,omitempty" xml:"interval_ids>string,omitempty"`
	// 100
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 100
	Pids []string `json:"pids,omitempty" xml:"pids>string,omitempty"`
}
