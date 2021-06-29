package alitrippoi

// AlitripPlatformPoiRawPoioutbypoiidsResult 
type AlitripPlatformPoiRawPoioutbypoiidsResult struct {
    // 返回poi详情
    Datas   []AlitripPlatformPoiRawPoioutbypoiidsData `json:"datas,omitempty" xml:"datas>alitrip_platform_poi_raw_poioutbypoiids_data,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 错误码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 总数(不可用)
    TotalRecords   int64 `json:"total_records,omitempty" xml:"total_records,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
