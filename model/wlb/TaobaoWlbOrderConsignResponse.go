package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流宝订单已发货通知接口 APIResponse
taobao.wlb.order.consign

如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
*/
type TaobaoWlbOrderConsignAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbOrderConsignResponse `json:"taobao_wlb_order_consign_response,omitempty"`
}

type TaobaoWlbOrderConsignResponse struct {

    // 修改时间
    ModifyTime   string `json:"modify_time,omitempty"`

}
