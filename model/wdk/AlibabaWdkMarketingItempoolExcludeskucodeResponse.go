package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池排除商品【品类优惠使用】 APIResponse
alibaba.wdk.marketing.itempool.excludeskucode

品类优惠新增排除池
*/
type AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolExcludeskucodeResponse
}

type AlibabaWdkMarketingItempoolExcludeskucodeResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_excludeskucode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品报名活动的返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
