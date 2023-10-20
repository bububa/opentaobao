package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse 根据码查询码信息（疫苗） API返回值
// alibaba.alihealth.drug.code.kyt.specia.vaccin.querycode
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponseModel is 根据码查询码信息（疫苗） 成功返回结果
type AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_specia_vaccin_querycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse
func GetAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse() *AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse 将 AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse(v *AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse.Put(v)
}
