package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse 通过企业名得到唯一标识(ref_ent_id)及企业ID(ent_id) API返回值
// alibaba.alihealth.drug.kyt.specia.vaccin.getentinfo
//
// 根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id)
type AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponseModel is 通过企业名得到唯一标识(ref_ent_id)及企业ID(ent_id) 成功返回结果
type AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_specia_vaccin_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse
func GetAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse() *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse {
	return poolAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse.Get().(*AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse 将 AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse(v *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse.Put(v)
}
