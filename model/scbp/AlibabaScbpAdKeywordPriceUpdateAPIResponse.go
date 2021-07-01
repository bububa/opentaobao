package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词改价 API返回值 
alibaba.scbp.ad.keyword.price.update

关键词改价
*/
type AlibabaScbpAdKeywordPriceUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordPriceUpdateAPIResponseModel
}

// 关键词改价 成功返回结果
type AlibabaScbpAdKeywordPriceUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_price_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 修改关键词价格是否成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
