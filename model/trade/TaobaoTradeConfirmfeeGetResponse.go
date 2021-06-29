package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取交易确认收货费用 APIResponse
taobao.trade.confirmfee.get

获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
*/
type TaobaoTradeConfirmfeeGetAPIResponse struct {
    model.CommonResponse
    TaobaoTradeConfirmfeeGetResponse
}

type TaobaoTradeConfirmfeeGetResponse struct {
    XMLName xml.Name `xml:"trade_confirmfee_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取到的交易确认收货费用
    
    TradeConfirmFee   *TradeConfirmFee `json:"trade_confirm_fee,omitempty" xml:"trade_confirm_fee,omitempty"`

    
}
