package xhotel

// TopAdsTripSvcQueryResult 结构体
type TopAdsTripSvcQueryResult struct {
	// topAdsHtlServiceDataList
	TopAdsHtlServiceDataList []Topadshtlservicedatalist `json:"top_ads_htl_service_data_list,omitempty" xml:"top_ads_htl_service_data_list>topadshtlservicedatalist,omitempty"`
	// count
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
