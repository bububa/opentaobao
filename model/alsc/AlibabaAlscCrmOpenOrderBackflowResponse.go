package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单回流接口 APIResponse
alibaba.alsc.crm.open.order.backflow

回流isv订单接口
*/
type AlibabaAlscCrmOpenOrderBackflowAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmOpenOrderBackflowResponse `json:"alibaba_alsc_crm_open_order_backflow_response,omitempty"`
}

type AlibabaAlscCrmOpenOrderBackflowResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
