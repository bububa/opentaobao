package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
溯源url透出 APIResponse
alibaba.wdk.trace.url.get

根据shopId和skuCode返回商品溯源url
*/
type AlibabaWdkTraceUrlGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkTraceUrlGetResponse `json:"alibaba_wdk_trace_url_get_response,omitempty"` 
    AlibabaWdkTraceUrlGetResponse
}

/* model for simplify = false
type AlibabaWdkTraceUrlGetResponse struct {

    // message
    
    Message   string `json:"message,omitempty"`
    

    // data
    
    Data   string `json:"data,omitempty"`
    

    // code
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type AlibabaWdkTraceUrlGetResponse struct {

    // message
    Message   string `json:"message,omitempty"`

    // data
    Data   string `json:"data,omitempty"`

    // code
    ReturnCode   string `json:"return_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
