package xhotel

// TopAdsSlrQueryResult 
type TopAdsSlrQueryResult struct {

    // 数据记录条数
    Count   int64 `json:"count,omitempty"`

    // adsSlrServiceDataList
    AdsSlrServiceDataList   []Adsslrservicedatalist `json:"ads_slr_service_data_list,omitempty"`

}