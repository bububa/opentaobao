package security

// UrlScanParamItem 
/* model for simplify = false
type UrlScanParamItem struct {

    // 需要扫描的url
    
    Data   string `json:"data,omitempty"`
    

    // 参数标记，用于识别返回结果对应的参数
    
    Flag   string `json:"flag,omitempty"`
    

}
*/

// UrlScanParamItem 
type UrlScanParamItem struct {

    // 需要扫描的url
    Data   string `json:"data,omitempty"`

    // 参数标记，用于识别返回结果对应的参数
    Flag   string `json:"flag,omitempty"`

}
