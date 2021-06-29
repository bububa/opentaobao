package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除买赠活动 API返回值 
alibaba.wdk.marketing.itembuygift.deleteactivity

删除买赠活动
*/
type AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItembuygiftDeleteactivityResponse
}

// 删除买赠活动 成功返回结果
type AlibabaWdkMarketingItembuygiftDeleteactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_deleteactivity_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 删除活动返回结果
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
