package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
移除买赠活动下的商品。【注意，此接口暂不支持并发！】 API返回值 
alibaba.wdk.marketing.itembuygift.removeitem

移除买赠活动下的商品。【注意，此接口暂不支持并发！】
*/
type AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItembuygiftRemoveitemAPIResponseModel
}

// 移除买赠活动下的商品。【注意，此接口暂不支持并发！】 成功返回结果
type AlibabaWdkMarketingItembuygiftRemoveitemAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_removeitem_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 移除商品返回结果
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
