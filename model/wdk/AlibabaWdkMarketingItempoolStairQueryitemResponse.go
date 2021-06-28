package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
换购商品查询 APIResponse
alibaba.wdk.marketing.itempool.stair.queryitem

换购商品查询
*/
type AlibabaWdkMarketingItempoolStairQueryitemAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itempool_stair_queryitem_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果
    
    Result   *MarketPageResult `json:"result,omitempty" xml:"