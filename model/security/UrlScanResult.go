package security

// UrlScanResult 
/* model for simplify = false
type UrlScanResult struct {

    // 请求标志唯一id
    
    EventId   string `json:"event_id,omitempty"`
    

    // 扫描详细结果
    
    UrlResultItemList  struct {
        UrlScanResultItem  []UrlScanResultItem `json:"url_scan_result_item,omitempty"`
    } `json:"url_result_item_list,omitempty"`
    

}
*/

// UrlScanResult 
type UrlScanResult struct {

    // 请求标志唯一id
    EventId   string `json:"event_id,omitempty"`

    // 扫描详细结果
    UrlResultItemList   []UrlScanResultItem `json:"url_result_item_list,omitempty"`

}
