package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
北京小蜜蜂配作业回传 APIResponse
alibaba.wdk.fulfill.dms.ebeecake.work.order.callback

北京小蜜蜂配作业回传。
*/
type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackResponse `json:"alibaba_wdk_fulfill_dms_ebeecake_work_order_callback_response,omitempty"`
}

type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackResponse struct {

    // 响应提示信息
    RespMessage   string `json:"resp_message,omitempty"`

    // 响应code
    RespCode   string `json:"resp_code,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
