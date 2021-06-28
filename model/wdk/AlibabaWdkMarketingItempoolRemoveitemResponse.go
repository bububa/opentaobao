package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
移除商品池里面的商品 APIResponse
alibaba.wdk.marketing.itempool.removeitem

移除商品池里面的商品
*/
type AlibabaWdkMarketingItempoolRemoveitemAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itempool_removeitem_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 移除商品返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"