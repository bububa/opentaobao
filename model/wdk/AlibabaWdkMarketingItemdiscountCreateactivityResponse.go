package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建商品特价活动 APIResponse
alibaba.wdk.marketing.itemdiscount.createactivity

创建商品特价活动
*/
type AlibabaWdkMarketingItemdiscountCreateactivityAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itemdiscount_createactivity_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建活动返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"