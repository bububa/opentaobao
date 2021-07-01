package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销退款接口 API返回值 
alibaba.wdk.pos.trade.reverse

轻pos品牌营销场景，商家调用退款接口
*/
type AlibabaWdkPosTradeReverseAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPosTradeReverseAPIResponseModel
}

// 轻pos品牌营销退款接口 成功返回结果
type AlibabaWdkPosTradeReverseAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_pos_trade_reverse_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 退款结果
    Result   *FastBuyPosReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}
