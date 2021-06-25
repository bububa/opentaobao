package wdk

// WdkMarketOpenResult 
type WdkMarketOpenResult struct {

    // 123
    Datas   []SyncActivitySkuResultBo `json:"datas,omitempty"`

    // 123123
    Success   bool `json:"success,omitempty"`

    // 123123
    ErrorCode   string `json:"error_code,omitempty"`

    // 123123
    Message   string `json:"message,omitempty"`

    // 活动数据
    ActivityList   []SyncActivityResultBo `json:"activity_list,omitempty"`

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // 版本信息
    Data   *SyncVersionBo `json:"data,omitempty"`

}
