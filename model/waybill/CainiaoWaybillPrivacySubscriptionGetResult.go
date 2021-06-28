package waybill

// CainiaoWaybillPrivacySubscriptionGetResult 
/* model for simplify = false
type CainiaoWaybillPrivacySubscriptionGetResult struct {

    // 错误code列表
    
    ErrorCodeList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_code_list,omitempty"`
    

    // 是否失败
    
    Failure   bool `json:"failure,omitempty"`
    

    // 第一个错误
    
    OneErrorInfo   string `json:"one_error_info,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 商家是否订购
    
    Subscription   bool `json:"subscription,omitempty"`
    

    // 系统自动生成
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误列表
    
    ErrorInfoList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_info_list,omitempty"`
    

    // 系统信息
    
    ObjectId   string `json:"object_id,omitempty"`
    

}
*/

// CainiaoWaybillPrivacySubscriptionGetResult 
type CainiaoWaybillPrivacySubscriptionGetResult struct {

    // 错误code列表
    ErrorCodeList   []string `json:"error_code_list,omitempty"`

    // 是否失败
    Failure   bool `json:"failure,omitempty"`

    // 第一个错误
    OneErrorInfo   string `json:"one_error_info,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 商家是否订购
    Subscription   bool `json:"subscription,omitempty"`

    // 系统自动生成
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误列表
    ErrorInfoList   []string `json:"error_info_list,omitempty"`

    // 系统信息
    ObjectId   string `json:"object_id,omitempty"`

}
