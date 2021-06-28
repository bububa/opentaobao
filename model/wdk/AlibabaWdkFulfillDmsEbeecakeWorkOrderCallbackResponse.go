package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
北京小蜜蜂配作业回传 APIResponse
alibaba.wdk.fulfill.dms.ebeecake.work.order.callback

北京小蜜蜂配作业回传。
*/
type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackResponse
}

type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_dms_ebeecake_work_order_callback_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应提示信息
    
    RespMessage   string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`

    
    // 响应code
    
    RespCode   string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
