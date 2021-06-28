package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销退款接口 APIResponse
alibaba.wdk.pos.trade.reverse

轻pos品牌营销场景，商家调用退款接口
*/
type AlibabaWdkPosTradeReverseAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPosTradeReverseResponse
}

type AlibabaWdkPosTradeReverseResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_pos_trade_reverse_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 退款结果
    
    Result   *FastBuyPosReverseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
