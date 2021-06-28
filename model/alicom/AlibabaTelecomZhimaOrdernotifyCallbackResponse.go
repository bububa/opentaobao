package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
阿里通信芝麻订单通知 APIResponse
alibaba.telecom.zhima.ordernotify.callback

商家通知阿里通信，芝麻订单状态，阿里通信侧进行代扣支付、发货并确认收货
*/
type AlibabaTelecomZhimaOrdernotifyCallbackAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTelecomZhimaOrdernotifyCallbackResponse `json:"alibaba_telecom_zhima_ordernotify_callback_response,omitempty"` 
    AlibabaTelecomZhimaOrdernotifyCallbackResponse
}

/* model for simplify = false
type AlibabaTelecomZhimaOrdernotifyCallbackResponse struct {

    // 出参对象
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaTelecomZhimaOrdernotifyCallbackResponse struct {

    // 出参对象
    Result   *CommonResult `json:"result,omitempty"`

}
