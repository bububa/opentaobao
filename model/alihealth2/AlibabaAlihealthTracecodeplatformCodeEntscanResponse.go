package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
药品商家扫码 APIResponse
alibaba.alihealth.tracecodeplatform.code.entscan

药品商家扫描药品监管码，只有该商家的药才返回
*/
type AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodeplatformCodeEntscanResponse
}

type AlibabaAlihealthTracecodeplatformCodeEntscanResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeplatform_code_entscan_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
