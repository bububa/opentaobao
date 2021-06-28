package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买赠活动下的商品 APIResponse
alibaba.wdk.marketing.itembuygift.queryitems

查询买赠活动下的商品
*/
type AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itembuygift_queryitems_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询返回结果
    
    Result   *MarketPageResult `json:"result,omitempty" xml:"