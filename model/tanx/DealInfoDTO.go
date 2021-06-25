package tanx

// DealInfoDTO 
type DealInfoDTO struct {

    // 11
    Addresses   []DicDTO `json:"addresses,omitempty"`

    // 100
    DealId   int64 `json:"deal_id,omitempty"`

    // 100
    Price   int64 `json:"price,omitempty"`

    // 100
    BeginTime   string `json:"begin_time,omitempty"`

    // 100
    EndTime   string `json:"end_time,omitempty"`

    // 100
    DealType   int64 `json:"deal_type,omitempty"`

    // 100
    SellerSiteNames   []String `json:"seller_site_names,omitempty"`

    // 100
    DspIds   []Number `json:"dsp_ids,omitempty"`

    // 100
    AdvertiserIds   []Number `json:"advertiser_ids,omitempty"`

    // 100
    IntervalIds   []String `json:"interval_ids,omitempty"`

    // 100
    Status   int64 `json:"status,omitempty"`

    // 100
    Pids   []String `json:"pids,omitempty"`

}
