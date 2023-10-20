package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesGetentinfoAPIResponse 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API返回值
// alibaba.alihealth.drug.kyt.wes.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaAlihealthDrugKytWesGetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesGetentinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesGetentinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesGetentinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesGetentinfoAPIResponseModel is 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 成功返回结果
type AlibabaAlihealthDrugKytWesGetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesGetentinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesGetentinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesGetentinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesGetentinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesGetentinfoAPIResponse
func GetAlibabaAlihealthDrugKytWesGetentinfoAPIResponse() *AlibabaAlihealthDrugKytWesGetentinfoAPIResponse {
	return poolAlibabaAlihealthDrugKytWesGetentinfoAPIResponse.Get().(*AlibabaAlihealthDrugKytWesGetentinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesGetentinfoAPIResponse 将 AlibabaAlihealthDrugKytWesGetentinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesGetentinfoAPIResponse(v *AlibabaAlihealthDrugKytWesGetentinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesGetentinfoAPIResponse.Put(v)
}
