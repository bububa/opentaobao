package security

// UrlScanParamList 
/* model for simplify = false
type UrlScanParamList struct {

    // 扫描参数列表
    
    UrlScanParamItems  struct {
        UrlScanParamItem  []UrlScanParamItem `json:"url_scan_param_item,omitempty"`
    } `json:"url_scan_param_items,omitempty"`
    

}
*/

// UrlScanParamList 
type UrlScanParamList struct {

    // 扫描参数列表
    UrlScanParamItems   []UrlScanParamItem `json:"url_scan_param_items,omitempty"`

}
