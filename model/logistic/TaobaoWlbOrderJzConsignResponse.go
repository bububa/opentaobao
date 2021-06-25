package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
家装发货接口 APIResponse
taobao.wlb.order.jz.consign

家装类订单使用该接口发货
*/
type TaobaoWlbOrderJzConsignAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbOrderJzConsignResponse `json:"taobao_wlb_order_jz_consign_response,omitempty"`
}

type TaobaoWlbOrderJzConsignResponse struct {

    // 错误码
    ResultErrorCode   string `json:"result_error_code,omitempty"`

    // 错误信息描述
    ResultErrorMsg   string `json:"result_error_msg,omitempty"`

    // 是否成功
    ResultSuccess   bool `json:"result_success,omitempty"`

}
