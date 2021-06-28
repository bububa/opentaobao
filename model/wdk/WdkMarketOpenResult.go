package wdk

// WdkMarketOpenResult 
/* model for simplify = false
type WdkMarketOpenResult struct {

    // 123
    
    Datas  struct {
        SyncActivitySkuResultBo  []SyncActivitySkuResultBo `json:"sync_activity_sku_result_bo,omitempty"`
    } `json:"datas,omitempty"`
    

    // 123123
    
    Success   bool `json:"success,omitempty"`
    

    // 123123
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 123123
    
    Message   string `json:"message,omitempty"`
    

    // 活动数据
    
    ActivityList  struct {
        SyncActivityResultBo  []SyncActivityResultBo `json:"sync_activity_result_bo,omitempty"`
    } `json:"activity_list,omitempty"`
    

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 版本信息
    
    Data  *struct {
        SyncVersionBo  *SyncVersionBo `json:"sync_version_bo,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

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
