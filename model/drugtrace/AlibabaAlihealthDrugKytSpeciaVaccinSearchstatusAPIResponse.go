package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse 疫苗企业上传单据后处理状态查询 API返回值
// alibaba.alihealth.drug.kyt.specia.vaccin.searchstatus
//
// 单据处理状态查询
type AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponseModel is 疫苗企业上传单据后处理状态查询 成功返回结果
type AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_specia_vaccin_searchstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse
func GetAlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse() *AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse {
	return poolAlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse.Get().(*AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse 将 AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse(v *AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse.Put(v)
}
