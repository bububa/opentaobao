package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码流向查询 APIResponse
alibaba.alihealth.drug.code.kyt.querycodeflow

追溯码流向查询
*/
type AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugCodeKytQuerycodeflowResponse
}

type AlibabaAlihealthDrugCodeKytQuerycodeflowResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_querycodeflow_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugCodeKytQuerycodeflowResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
