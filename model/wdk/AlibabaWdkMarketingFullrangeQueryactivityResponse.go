package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动查询活动 API返回值 
alibaba.wdk.marketing.fullrange.queryactivity

全场活动查询活动
*/
type AlibabaWdkMarketingFullrangeQueryactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingFullrangeQueryactivityResponse
}

// 全场活动查询活动 成功返回结果
type AlibabaWdkMarketingFullrangeQueryactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_queryactivity_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询返回结果
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
