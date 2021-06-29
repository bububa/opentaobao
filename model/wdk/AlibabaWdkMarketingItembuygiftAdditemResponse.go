package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加买赠活动商品。【注意，此接口暂不支持并发！】 API返回值 
alibaba.wdk.marketing.itembuygift.additem

增加买赠活动商品。【注意，此接口暂不支持并发！】
*/
type AlibabaWdkMarketingItembuygiftAdditemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItembuygiftAdditemResponse
}

// 增加买赠活动商品。【注意，此接口暂不支持并发！】 成功返回结果
type AlibabaWdkMarketingItembuygiftAdditemResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_additem_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商品报名活动的返回结果
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
