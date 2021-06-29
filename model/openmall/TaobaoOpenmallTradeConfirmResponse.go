package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认收货 APIResponse
taobao.openmall.trade.confirm

确认订单收货
*/
type TaobaoOpenmallTradeConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallTradeConfirmResponse
}

type TaobaoOpenmallTradeConfirmResponse struct {
    XMLName xml.Name `xml:"openmall_trade_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}
