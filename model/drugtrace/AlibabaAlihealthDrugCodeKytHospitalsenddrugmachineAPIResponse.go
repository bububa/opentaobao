package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse 医院发药机 API返回值
// alibaba.alihealth.drug.code.kyt.hospitalsenddrugmachine
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
// 提供给医院发药机使用
type AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponseModel is 医院发药机 成功返回结果
type AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_hospitalsenddrugmachine_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse
func GetAlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse() *AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse 将 AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse(v *AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse.Put(v)
}
