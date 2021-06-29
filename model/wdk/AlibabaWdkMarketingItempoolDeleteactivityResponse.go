package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动 APIResponse
alibaba.wdk.marketing.itempool.deleteactivity

删除商品池活动
*/
type AlibabaWdkMarketingItempoolDeleteactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolDeleteactivityResponse
}

type AlibabaWdkMarketingItempoolDeleteactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_deleteactivity_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 删除活动返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
