package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
是否为新支付订单 APIResponse
alibaba.mos.onsite.trade.isnewpayorder

退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单
*/
type AlibabaMosOnsiteTradeIsnewpayorderAPIResponse struct {
    model.CommonResponse
    AlibabaMosOnsiteTradeIsnewpayorderResponse
}

type AlibabaMosOnsiteTradeIsnewpayorderResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_onsite_trade_isnewpayorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaMosOnsiteTradeIsnewpayorderResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
