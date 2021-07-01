package xhotel

// TopAdsHtlDataQueryResult 结构体
type TopAdsHtlDataQueryResult struct {
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// topAdsHtlServiceDataList
	TopAdsHtlServiceDataList []Topadshtlservicedatalist `json:"top_ads_htl_service_data_list,omitempty" xml:"top_ads_htl_service_data_list>topadshtlservicedatalist,omitempty"`
}
