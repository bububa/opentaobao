package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse 根据码获取码信息接口-医保 API返回值
// alibaba.alihealth.drug.code.list.code.medical.insurance
//
// 服务描述
// 医保鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
// 与验证鉴核流程；
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
// 若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
// 情况下，需要分多次调用该接口。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponseModel is 根据码获取码信息接口-医保 成功返回结果
type AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_list_code_medical_insurance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse
func GetAlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse() *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse {
	return poolAlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse.Get().(*AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse 将 AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse(v *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIResponse.Put(v)
}
