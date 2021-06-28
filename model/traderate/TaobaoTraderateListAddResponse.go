package traderate

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
针对父子订单新增批量评价 APIResponse
taobao.traderate.list.add

针对父子订单新增批量评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价</font>)
*/
type TaobaoTraderateListAddAPIResponse struct {
    model.CommonResponse
    TaobaoTraderateListAddResponse
}

type TaobaoTraderateListAddResponse struct {
    XMLName xml.Name `xml:"traderate_list_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的评论的信息，仅返回tid和created字段
    
    TradeRate   *TradeRateRequest `json:"trade_rate,omitempty" xml:"trade_rate,omitempty"`

    
}
