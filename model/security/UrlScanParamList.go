package security

// UrlScanParamList 
type UrlScanParamList struct {
    // 扫描参数列表
    UrlScanParamItems   []UrlScanParamItem `json:"url_scan_param_items,omitempty" xml:"url_scan_param_items>url_scan_param_item,omitempty"`
}
