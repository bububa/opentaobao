package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询扫码信息 APIResponse
alibaba.alihealth.drugcode.scan

查询扫码信息
*/
type AlibabaAlihealthDrugcodeScanAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeScanResponse
}

type AlibabaAlihealthDrugcodeScanResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_scan_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
