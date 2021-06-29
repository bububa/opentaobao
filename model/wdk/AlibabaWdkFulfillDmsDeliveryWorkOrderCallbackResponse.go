package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
末端配配送作业回传 API返回值 
alibaba.wdk.fulfill.dms.delivery.work.order.callback

末端配配送作业回传。
*/
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackResponse
}

// 末端配配送作业回传 成功返回结果
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_dms_delivery_work_order_callback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应提示信息
    RespMessage   string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`
    // 返回码： SUCCESS:作业批次接收成功 SYSTEM_ERROR :系统异常(指令可重发) PARAM_ERROR :参数错误(指令不可重发，监控报警) BUSINESS_ERROR:业务异常(指令不可重发，监控报警)
    RespCode   string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
    // 是否成功：true 调用成功； false 调用失败
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
