package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单码关联关系查询，通过一个码查询这个码下的所有子码 APIResponse
alibaba.alihealth.drug.kyt.singlerelation

单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
*/
type AlibabaAlihealthDrugKytSinglerelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytSinglerelationResponse
}

type AlibabaAlihealthDrugKytSinglerelationResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_singlerelation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytSinglerelationResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
