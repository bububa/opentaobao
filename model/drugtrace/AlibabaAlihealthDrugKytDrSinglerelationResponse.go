package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
多融单码关联关系查询 APIResponse
alibaba.alihealth.drug.kyt.dr.singlerelation

单码关联关系查询
*/
type AlibabaAlihealthDrugKytDrSinglerelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrSinglerelationResponse
}

type AlibabaAlihealthDrugKytDrSinglerelationResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_singlerelation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytDrSinglerelationResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
