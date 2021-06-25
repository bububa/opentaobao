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
    Response *AlibabaWdkUmsOutboundProcessGetResponse `json:"alibaba_wdk_ums_outbound_process_get_response,omitempty"`
}

type AlibabaWdkUmsOutboundProcessGetResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
