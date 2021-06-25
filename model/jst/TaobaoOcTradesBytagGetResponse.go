package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
标签查询订单 APIResponse
taobao.oc.trades.bytag.get

根据标签查询订单编号
*/
type TaobaoOcTradesBytagGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOcTradesBytagGetResponse `json:"taobao_oc_trades_bytag_get_response,omitempty"`
}

type TaobaoOcTradesBytagGetResponse struct {

    // 打了该标签的订单编号列表
    Tids   []Number `json:"tids,omitempty"`

    // 总数
    Totals   int64 `json:"totals,omitempty"`

}
