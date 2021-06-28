package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池删除商品 APIResponse
alibaba.wdk.marketing.itempool.item.remove.async

新模型下删除商品
*/
type AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itempool_item_remove_async_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Result   *MarketResult `json:"result,omitempty" xml:"