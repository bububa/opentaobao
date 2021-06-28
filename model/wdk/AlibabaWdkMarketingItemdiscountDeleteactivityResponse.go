package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品特价活动 APIResponse
alibaba.wdk.marketing.itemdiscount.deleteactivity

删除商品特价活动
*/
type AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItemdiscountDeleteactivityResponse
}

type AlibabaWdkMarketingItemdiscountDeleteactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_deleteactivity_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除活动返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
