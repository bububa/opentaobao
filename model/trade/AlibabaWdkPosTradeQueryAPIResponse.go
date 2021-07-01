package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销查询接口 API返回值 
alibaba.wdk.pos.trade.query

轻pos品牌营销场景，外部商家查询营销信息
*/
type AlibabaWdkPosTradeQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPosTradeQueryAPIResponseModel
}

// 轻pos品牌营销查询接口 成功返回结果
type AlibabaWdkPosTradeQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_pos_trade_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询返回结果
    Result   *FastBuyPosQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
