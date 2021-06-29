package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
药品是否赋码 API返回值 
alibaba.alihealth.drug.kyt.drugcodes

药品是否赋码
*/
type AlibabaAlihealthDrugKytDrugcodesAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrugcodesResponse
}

// 药品是否赋码 成功返回结果
type AlibabaAlihealthDrugKytDrugcodesResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_drugcodes_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 监控宝推送网站监控信息，返回结果
    Result   *AlibabaAlihealthDrugKytDrugcodesResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
