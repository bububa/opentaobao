package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除换购活动商品 APIResponse
alibaba.wdk.marketing.itempool.stair.removeitem

删除换购商品
*/
type AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolStairRemoveitemResponse
}

type AlibabaWdkMarketingItempoolStairRemoveitemResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_stair_removeitem_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
