package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建OC订单 APIResponse
taobao.oc.order.create

创建OC订单接口
*/
type TaobaoOcOrderCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOcOrderCreateResponse `json:"taobao_oc_order_create_response,omitempty"`
}

type TaobaoOcOrderCreateResponse struct {

    // 表示是否执行成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 执行失败原因
    ErrorDescription   string `json:"error_description,omitempty"`

    // OcOrder
    OcOrder   *OcOrder `json:"oc_order,omitempty"`

}
