package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
移除报名的商品 APIResponse
alibaba.wdk.marketing.itemdiscount.removeitem

将报名特价活动的商品从特价活动中移除
*/
type AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itemdiscount_removeitem_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 移除商品返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"