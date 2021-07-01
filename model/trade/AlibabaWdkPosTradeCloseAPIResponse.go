package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销关单接口 API返回值 
alibaba.wdk.pos.trade.close

轻pos品牌营销场景，提供关单接口给外部商家
*/
type AlibabaWdkPosTradeCloseAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPosTradeCloseAPIResponseModel
}

// 轻pos品牌营销关单接口 成功返回结果
type AlibabaWdkPosTradeCloseAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_pos_trade_close_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 关单结果
    Result   *FastBuyPosCloseResult `json:"result,omitempty" xml:"result,omitempty"`
}
