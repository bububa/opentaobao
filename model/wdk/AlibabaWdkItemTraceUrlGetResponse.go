package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据shopId和skuCode返回商品静态溯源url APIResponse
alibaba.wdk.item.trace.url.get

根据shopId和skuCode返回商品静态溯源url
*/
type AlibabaWdkItemTraceUrlGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_item_trace_url_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // message
    
    Message   string `json:"message,omitempty" xml:"