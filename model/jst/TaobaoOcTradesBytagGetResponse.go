package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
标签查询订单 APIResponse
taobao.oc.trades.bytag.get

根据标签查询订单编号
*/
type TaobaoOcTradesBytagGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"oc_trades_bytag_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 打了该标签的订单编号列表
    
    Tids   []int64 `json:"tids,omitempty" xml:"