package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售+平台支付订单 APIResponse
alibaba.nlife.b2c.trade.pay

零售+平台支付接口，外部商户调用此接口告知零售+支付结果，保持订单状态同步
*/
type AlibabaNlifeB2cTradePayAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeB2cTradePayResponse
}

type AlibabaNlifeB2cTradePayResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_pay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 扩展参数
    
    ExtendParams   string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`

    
    // gmtPayment
    
    GmtPayment   string `json:"gmt_payment,omitempty" xml:"gmt_payment,omitempty"`

    
}
