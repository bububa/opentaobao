package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买赠活动下的商品 API返回值 
alibaba.wdk.marketing.itembuygift.queryitems

查询买赠活动下的商品
*/
type AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItembuygiftQueryitemsResponse
}

// 查询买赠活动下的商品 成功返回结果
type AlibabaWdkMarketingItembuygiftQueryitemsResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_queryitems_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询返回结果
    Result   *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
