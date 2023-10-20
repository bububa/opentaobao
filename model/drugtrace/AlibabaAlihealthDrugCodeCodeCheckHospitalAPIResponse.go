package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse 码核查状态同步-医院 API返回值
// alibaba.alihealth.drug.code.code.check.hospital
//
// 码核查状态同步-医院
type AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponseModel is 码核查状态同步-医院 成功返回结果
type AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_code_check_hospital_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse
func GetAlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse() *AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse {
	return poolAlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse.Get().(*AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse 将 AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse(v *AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse.Put(v)
}
