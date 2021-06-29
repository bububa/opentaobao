package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销支付接口 APIResponse
alibaba.wdk.pos.trade.pay

轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
*/
type AlibabaWdkPosTradePayAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPosTradePayResponse
}

type AlibabaWdkPosTradePayResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_pos_trade_pay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 支付结果
    
    Result   *FastBuyPosPayResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
