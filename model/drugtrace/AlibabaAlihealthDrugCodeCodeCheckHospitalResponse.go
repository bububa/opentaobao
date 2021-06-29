package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码核查状态同步-医院 API返回值 
alibaba.alihealth.drug.code.code.check.hospital

码核查状态同步-医院
*/
type AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugCodeCodeCheckHospitalResponse
}

// 码核查状态同步-医院 成功返回结果
type AlibabaAlihealthDrugCodeCodeCheckHospitalResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_code_code_check_hospital_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 监控宝推送网站监控信息，返回结果
    Result   *AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
