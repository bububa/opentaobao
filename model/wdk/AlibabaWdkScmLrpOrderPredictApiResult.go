package wdk

// AlibabaWdkScmLrpOrderPredictApiResult 
/* model for simplify = false
type AlibabaWdkScmLrpOrderPredictApiResult struct {

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误消息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 单量预测结果列表
    
    PredictList  struct {
        OrderPredict  []OrderPredict `json:"order_predict,omitempty"`
    } `json:"predict_list,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkScmLrpOrderPredictApiResult 
type AlibabaWdkScmLrpOrderPredictApiResult struct {

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误消息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 单量预测结果列表
    PredictList   []OrderPredict `json:"predict_list,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
