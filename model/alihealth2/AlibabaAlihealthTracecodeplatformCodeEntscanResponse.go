package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
药品商家扫码 API返回值 
alibaba.alihealth.tracecodeplatform.code.entscan

药品商家扫描药品监管码，只有该商家的药才返回
*/
type AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodeplatformCodeEntscanResponse
}

// 药品商家扫码 成功返回结果
type AlibabaAlihealthTracecodeplatformCodeEntscanResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeplatform_code_entscan_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
