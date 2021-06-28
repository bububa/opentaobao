package tanx

// DealInfoDTO 
/* model for simplify = false
type DealInfoDTO struct {

    // 11
    
    Addresses  struct {
        DicDTO  []DicDTO `json:"dic_dto,omitempty"`
    } `json:"addresses,omitempty"`
    

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
    
    SellerSiteNames  struct {
        String  []string `json:"string,omitempty"`
    } `json:"seller_site_names,omitempty"`
    

    // 100
    
    DspIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"dsp_ids,omitempty"`
    

    // 100
    
    AdvertiserIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"advertiser_ids,omitempty"`
    

    // 100
    
    IntervalIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"interval_ids,omitempty"`
    

    // 100
    
    Status   int64 `json:"status,omitempty"`
    

    // 100
    
    Pids  struct {
        String  []string `json:"string,omitempty"`
    } `json:"pids,omitempty"`
    

}
*/

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
    SellerSiteNames   []string `json:"seller_site_names,omitempty"`

    // 100
    DspIds   []int64 `json:"dsp_ids,omitempty"`

    // 100
    AdvertiserIds   []int64 `json:"advertiser_ids,omitempty"`

    // 100
    IntervalIds   []string `json:"interval_ids,omitempty"`

    // 100
    Status   int64 `json:"status,omitempty"`

    // 100
    Pids   []string `json:"pids,omitempty"`

}
