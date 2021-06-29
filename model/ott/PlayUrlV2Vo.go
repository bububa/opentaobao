package ott

// PlayUrlV2Vo 
type PlayUrlV2Vo struct {
    // 订单状态
    OrderStatus   int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // ytid
    Ytid   int64 `json:"ytid,omitempty" xml:"ytid,omitempty"`
    // hlsContent
    HlsContent   string `json:"hls_content,omitempty" xml:"hls_content,omitempty"`
    // hlsContentUrl
    HlsContentUrl   string `json:"hls_content_url,omitempty" xml:"hls_content_url,omitempty"`
    // dashContent
    DashContent   string `json:"dash_content,omitempty" xml:"dash_content,omitempty"`
    // drmToken
    DrmToken   string `json:"drm_token,omitempty" xml:"drm_token,omitempty"`
    // errCode
    ErrCode   int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // duration
    Duration   int64 `json:"duration,omitempty" xml:"duration,omitempty"`
    // programId
    ProgramId   int64 `json:"program_id,omitempty" xml:"program_id,omitempty"`
    // productType
    ProductType   int64 `json:"product_type,omitempty" xml:"product_type,omitempty"`
    // sourceInfo
    SourceInfo   *SourceInfo `json:"source_info,omitempty" xml:"source_info,omitempty"`
    // httpDns
    HttpDns   *HttpDns `json:"http_dns,omitempty" xml:"http_dns,omitempty"`
    // trial
    Trial   bool `json:"trial,omitempty" xml:"trial,omitempty"`
    // tokenValid
    TokenValid   bool `json:"token_valid,omitempty" xml:"token_valid,omitempty"`
    // free
    Free   bool `json:"free,omitempty" xml:"free,omitempty"`
    // overDeviceLimit
    OverDeviceLimit   bool `json:"over_device_limit,omitempty" xml:"over_device_limit,omitempty"`
    // live
    Live   bool `json:"live,omitempty" xml:"live,omitempty"`
    // 片头
    StartTime   int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 片尾
    EndTime   int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // vr
    Vr   bool `json:"vr,omitempty" xml:"vr,omitempty"`
}
