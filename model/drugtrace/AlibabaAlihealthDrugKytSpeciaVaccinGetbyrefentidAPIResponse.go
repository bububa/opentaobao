package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse 根据企业唯一标识查看企业详情 API返回值
// alibaba.alihealth.drug.kyt.specia.vaccin.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponseModel is 根据企业唯一标识查看企业详情 成功返回结果
type AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_specia_vaccin_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse
func GetAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse() *AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse {
	return poolAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse.Get().(*AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse 将 AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse(v *AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse.Put(v)
}
