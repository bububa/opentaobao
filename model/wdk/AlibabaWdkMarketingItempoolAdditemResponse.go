package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品池里面的商品 APIResponse
alibaba.wdk.marketing.itempool.additem

增加商品池里面的商品
*/
type AlibabaWdkMarketingItempoolAdditemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolAdditemResponse
}

type AlibabaWdkMarketingItempoolAdditemResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_additem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品报名活动的返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
