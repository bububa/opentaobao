package wdk

// WdkMarketOpenResult 
type WdkMarketOpenResult struct {
    // 123
    Datas   []SyncActivitySkuResultBo `json:"datas,omitempty" xml:"datas>sync_activity_sku_result_bo,omitempty"`
    // 123123
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 123123
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 123123
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 活动数据
    ActivityList   []SyncActivityResultBo `json:"activity_list,omitempty" xml:"activity_list>sync_activity_result_bo,omitempty"`
    // 错误编码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 版本信息
    Data   *SyncVersionBo `json:"data,omitempty" xml:"data,omitempty"`
}
