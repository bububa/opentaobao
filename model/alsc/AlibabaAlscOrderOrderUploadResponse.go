package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单回流 APIResponse
alibaba.alsc.order.order.upload

第三方订单回流
*/
type AlibabaAlscOrderOrderUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_order_order_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"