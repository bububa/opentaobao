package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
末端配配送作业回传 APIResponse
alibaba.wdk.fulfill.dms.delivery.work.order.callback

末端配配送作业回传。
*/
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackResponse `json:"alibaba_wdk_fulfill_dms_delivery_work_order_callback_response,omitempty"`
}

type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackResponse struct {

    // 响应提示信息
    RespMessage   string `json:"resp_message,omitempty"`

    // 返回码： SUCCESS:作业批次接收成功 SYSTEM_ERROR :系统异常(指令可重发) PARAM_ERROR :参数错误(指令不可重发，监控报警) BUSINESS_ERROR:业务异常(指令不可重发，监控报警)
    RespCode   string `json:"resp_code,omitempty"`

    // 是否成功：true 调用成功； false 调用失败
    IsSuccess   bool `json:"is_success,omitempty"`

}
