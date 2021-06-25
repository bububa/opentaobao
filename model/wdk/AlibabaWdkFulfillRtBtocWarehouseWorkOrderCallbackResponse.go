package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
大润发B2C仓作业单回传接口 APIResponse
alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback

大润发B2C仓作业单回传接口
*/
type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackResponse `json:"alibaba_wdk_fulfill_rt_btoc_warehouse_work_order_callback_response,omitempty"`
}

type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 响应提示信息
    RespMessage   string `json:"resp_message,omitempty"`

    // 响应code
    RespCode   string `json:"resp_code,omitempty"`

}
