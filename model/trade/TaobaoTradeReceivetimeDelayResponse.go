package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
延长交易收货时间 APIResponse
taobao.trade.receivetime.delay

延长交易收货时间
*/
type TaobaoTradeReceivetimeDelayAPIResponse struct {
    model.CommonResponse
    TaobaoTradeReceivetimeDelayResponse
}

type TaobaoTradeReceivetimeDelayResponse struct {
    XMLName xml.Name `xml:"trade_receivetime_delay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新后的交易数据，只包括tid和modified两个字段。
    
    Trade   *Trade `json:"trade,omitempty" xml:"trade,omitempty"`

    
}
