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
    AlibabaWdkItemTraceUrlGetResponse
}

type AlibabaWdkItemTraceUrlGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_trace_url_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // data
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // code
    
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
