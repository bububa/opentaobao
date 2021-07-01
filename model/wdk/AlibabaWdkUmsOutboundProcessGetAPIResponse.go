package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出库业务UMS异步处理结果返回 API返回值 
alibaba.wdk.ums.outbound.process.get

出库业务UMS异步处理结果返回
*/
type AlibabaWdkUmsOutboundProcessGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsOutboundProcessGetAPIResponseModel
}

// 出库业务UMS异步处理结果返回 成功返回结果
type AlibabaWdkUmsOutboundProcessGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_outbound_process_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
