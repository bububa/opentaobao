package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
出库业务UMS异步处理结果返回 APIResponse
alibaba.wdk.ums.outbound.process.get

出库业务UMS异步处理结果返回
*/
type AlibabaWdkUmsOutboundProcessGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkUmsOutboundProcessGetResponse `json:"alibaba_wdk_ums_outbound_process_get_response,omitempty"` 
    AlibabaWdkUmsOutboundProcessGetResponse
}

/* model for simplify = false
type AlibabaWdkUmsOutboundProcessGetResponse struct {

    // result
    
    Result  *struct {
        UtmsResult  *UtmsResult `json:"utms_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkUmsOutboundProcessGetResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
