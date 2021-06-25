package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据shopId和skuCode返回商品静态溯源url APIResponse
alibaba.wdk.item.trace.url.get

根据shopId和skuCode返回商品静态溯源url
*/
type AlibabaWdkItemTraceUrlGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkItemTraceUrlGetResponse `json:"alibaba_wdk_item_trace_url_get_response,omitempty"`
}

type AlibabaWdkItemTraceUrlGetResponse struct {

    // message
    Message   string `json:"message,omitempty"`

    // data
    Data   string `json:"data,omitempty"`

    // code
    ReturnCode   string `json:"return_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
