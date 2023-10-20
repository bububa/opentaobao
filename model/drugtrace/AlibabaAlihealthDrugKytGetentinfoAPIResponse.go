package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetentinfoAPIResponse 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API返回值
// alibaba.alihealth.drug.kyt.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaAlihealthDrugKytGetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytGetentinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetentinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytGetentinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytGetentinfoAPIResponseModel is 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 成功返回结果
type AlibabaAlihealthDrugKytGetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetentinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytGetentinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetentinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytGetentinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetentinfoAPIResponse
func GetAlibabaAlihealthDrugKytGetentinfoAPIResponse() *AlibabaAlihealthDrugKytGetentinfoAPIResponse {
	return poolAlibabaAlihealthDrugKytGetentinfoAPIResponse.Get().(*AlibabaAlihealthDrugKytGetentinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytGetentinfoAPIResponse 将 AlibabaAlihealthDrugKytGetentinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetentinfoAPIResponse(v *AlibabaAlihealthDrugKytGetentinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetentinfoAPIResponse.Put(v)
}
