package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单查询订单标签 APIResponse
taobao.oc.tradetags.get

根据订单查询订单标签。<br/>
返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/>
*/
type TaobaoOcTradetagsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"oc_tradetags_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    TradeTags   []TradeTagRelationDo `json:"trade_tags,omitempty" xml:"