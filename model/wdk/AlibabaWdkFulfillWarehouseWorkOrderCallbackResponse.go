package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
标准化仓作业单回传接口 APIResponse
alibaba.wdk.fulfill.warehouse.work.order.callback

标准化仓作业单回传接口
*/
type AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillWarehouseWorkOrderCallbackResponse `json:"alibaba_wdk_fulfill_warehouse_work_order_callback_response,omitempty"`
}

type AlibabaWdkFulfillWarehouseWorkOrderCallbackResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 响应提示信息
    RespMessage   string `json:"resp_message,omitempty"`

    // 响应code
    RespCode   string `json:"resp_code,omitempty"`

}
