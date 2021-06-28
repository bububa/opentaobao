package security

// AppInfoBatch 
type AppInfoBatch struct {

    // APP应用类型 1-android 2-ios(暂不支持)
    
    AppOsType   int64 `json:"app_os_type,omitempty" xml:"app_os_type,omitempty"`
    

    // 回调地址,dataType=4时必填,用于处理完成后反向通知,通知为GET请求,请求格式:  callbackUrl+"?itemId=xxx&taskStatus=1"
    
    CallbackUrl   string `json:"callback_url,omitempty" xml:"callback_url,omitempty"`
    

    // APP数据类型 3-Batch MD5 4-Batch URL(暂不支持)
    
    DataType   int64 `json:"data_type,omitempty" xml:"data_type,omitempty"`
    

    // 需要扫描的应用的具体信息列表
    
    ScanInfos   []AppInfoBatchItem `json:"scan_infos,omitempty" xml:"scan_infos,omitempty"`
    

}
