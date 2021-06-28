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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itempool_additem_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品报名活动的返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"