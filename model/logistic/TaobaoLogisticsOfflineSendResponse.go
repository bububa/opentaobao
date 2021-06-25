package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
自己联系物流（线下物流）发货 APIResponse
taobao.logistics.offline.send

用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
*/
type TaobaoLogisticsOfflineSendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLogisticsOfflineSendResponse `json:"taobao_logistics_offline_send_response,omitempty"`
}

type TaobaoLogisticsOfflineSendResponse struct {

    // 自己联系的调用结果
    Shipping   *Shipping `json:"shipping,omitempty"`

}
