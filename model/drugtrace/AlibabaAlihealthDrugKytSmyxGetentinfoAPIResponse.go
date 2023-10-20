package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse 查企业标识信息 API返回值
// alibaba.alihealth.drug.kyt.smyx.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponseModel is 查企业标识信息 成功返回结果
type AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_smyx_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSmyxGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse
func GetAlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse() *AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse {
	return poolAlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse.Get().(*AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse 将 AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse(v *AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse.Put(v)
}
