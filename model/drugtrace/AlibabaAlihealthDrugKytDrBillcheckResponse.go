package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗追溯验证 APIResponse
alibaba.alihealth.drug.kyt.dr.billcheck

各级疾控在入库完成后，需要做追溯信息验证
*/
type AlibabaAlihealthDrugKytDrBillcheckAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrBillcheckResponse
}

type AlibabaAlihealthDrugKytDrBillcheckResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_billcheck_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytDrBillcheckResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
