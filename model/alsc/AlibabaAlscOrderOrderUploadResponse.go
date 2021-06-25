package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单回流 APIResponse
alibaba.alsc.order.order.upload

第三方订单回流
*/
type AlibabaAlscOrderOrderUploadAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscOrderOrderUploadResponse `json:"alibaba_alsc_order_order_upload_response,omitempty"`
}

type AlibabaAlscOrderOrderUploadResponse struct {

    // 返回包装类
    Result   *BaseResult `json:"result,omitempty"`

}
