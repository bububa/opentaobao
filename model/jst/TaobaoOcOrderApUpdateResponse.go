package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按OC订单分账 APIResponse
taobao.oc.order.ap.update

对OC订单执行分账操作
*/
type TaobaoOcOrderApUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOcOrderApUpdateResponse `json:"taobao_oc_order_ap_update_response,omitempty"`
}

type TaobaoOcOrderApUpdateResponse struct {

    // 描述操作执行是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // OC订单id
    OcOrderId   int64 `json:"oc_order_id,omitempty"`

    // 执行失败时候的错误描述信息
    ErrorDescription   string `json:"error_description,omitempty"`

}
