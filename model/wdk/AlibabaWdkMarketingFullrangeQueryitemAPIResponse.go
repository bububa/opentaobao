package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动查询换购品 API返回值 
alibaba.wdk.marketing.fullrange.queryitem

全场活动查询换购品
*/
type AlibabaWdkMarketingFullrangeQueryitemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingFullrangeQueryitemAPIResponseModel
}

// 全场活动查询换购品 成功返回结果
type AlibabaWdkMarketingFullrangeQueryitemAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_queryitem_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果
    Result   *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
