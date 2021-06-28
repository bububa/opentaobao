package xhotel

// TopAdsHtlDataQueryResult 
/* model for simplify = false
type TopAdsHtlDataQueryResult struct {

    // 数量
    
    Count   int64 `json:"count,omitempty"`
    

    // topAdsHtlServiceDataList
    
    TopAdsHtlServiceDataList  struct {
        Topadshtlservicedatalist  []Topadshtlservicedatalist `json:"topadshtlservicedatalist,omitempty"`
    } `json:"top_ads_htl_service_data_list,omitempty"`
    

}
*/

// TopAdsHtlDataQueryResult 
type TopAdsHtlDataQueryResult struct {

    // 数量
    Count   int64 `json:"count,omitempty"`

    // topAdsHtlServiceDataList
    TopAdsHtlServiceDataList   []Topadshtlservicedatalist `json:"top_ads_htl_service_data_list,omitempty"`

}
