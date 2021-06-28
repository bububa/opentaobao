package xhotel

// TopAdsTripSvcQueryResult 
/* model for simplify = false
type TopAdsTripSvcQueryResult struct {

    // topAdsHtlServiceDataList
    
    TopAdsHtlServiceDataList  struct {
        Topadshtlservicedatalist  []Topadshtlservicedatalist `json:"topadshtlservicedatalist,omitempty"`
    } `json:"top_ads_htl_service_data_list,omitempty"`
    

    // count
    
    Count   int64 `json:"count,omitempty"`
    

}
*/

// TopAdsTripSvcQueryResult 
type TopAdsTripSvcQueryResult struct {

    // topAdsHtlServiceDataList
    TopAdsHtlServiceDataList   []Topadshtlservicedatalist `json:"top_ads_htl_service_data_list,omitempty"`

    // count
    Count   int64 `json:"count,omitempty"`

}
