package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
换购商品查询 APIResponse
alibaba.wdk.marketing.itempool.stair.queryitem

换购商品查询
*/
type AlibabaWdkMarketingItempoolStairQueryitemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolStairQueryitemResponse
}

type AlibabaWdkMarketingItempoolStairQueryitemResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_stair_queryitem_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果
    
    Result   *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
