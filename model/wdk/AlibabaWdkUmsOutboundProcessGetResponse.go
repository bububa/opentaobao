package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出库业务UMS异步处理结果返回 APIResponse
alibaba.wdk.ums.outbound.process.get

出库业务UMS异步处理结果返回
*/
type AlibabaWdkUmsOutboundProcessGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsOutboundProcessGetResponse
}

type AlibabaWdkUmsOutboundProcessGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_outbound_process_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
