package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查找特价活动 API返回值 
alibaba.wdk.marketing.itemdiscount.queryactivity

查找特价活动
*/
type AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItemdiscountQueryactivityResponse
}

// 查找特价活动 成功返回结果
type AlibabaWdkMarketingItemdiscountQueryactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_queryactivity_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询特价活动返回结果
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
