package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查找特价活动 APIResponse
alibaba.wdk.marketing.itemdiscount.queryactivity

查找特价活动
*/
type AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itemdiscount_queryactivity_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询特价活动返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"