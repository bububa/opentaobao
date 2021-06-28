package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查找商品池活动 APIResponse
alibaba.wdk.marketing.itempool.queryactivity

查找商品池活动
*/
type AlibabaWdkMarketingItempoolQueryactivityAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itempool_queryactivity_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"