package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据追溯码获取69码 API返回值 
alibaba.alihealth.drug.getbarcode.bytraccode

根据追溯码获取69码
*/
type AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponseModel
}

// 根据追溯码获取69码 成功返回结果
type AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_getbarcode_bytraccode_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
