package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
每日优鲜仓作业单回传接口 APIResponse
alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback

家乐福仓作业单回传接口
*/
type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackResponse `json:"alibaba_wdk_fulfill_missfresh_warehouse_work_order_callback_response,omitempty"`
}

type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 响应提示信息
    RespMessage   string `json:"resp_message,omitempty"`

    // 响应code
    RespCode   string `json:"resp_code,omitempty"`

}
