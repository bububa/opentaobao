package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询特价商品 APIResponse
alibaba.wdk.marketing.itemdiscount.queryitems

查询参加特价活动的商品优惠详情
*/
type AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itemdiscount_queryitems_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询返回结果
    
    Result   *MarketPageResult `json:"result,omitempty" xml:"