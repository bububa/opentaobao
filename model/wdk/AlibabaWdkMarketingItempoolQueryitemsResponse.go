package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品池活动下的商品 APIResponse
alibaba.wdk.marketing.itempool.queryitems

查询商品池活动下面的商品
*/
type AlibabaWdkMarketingItempoolQueryitemsAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolQueryitemsResponse
}

type AlibabaWdkMarketingItempoolQueryitemsResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_queryitems_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询返回结果
    
    Result   *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
