package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSmyxListpartsAPIResponse 药店查询往来单位 API返回值
// alibaba.alihealth.drug.kyt.smyx.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugKytSmyxListpartsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSmyxListpartsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSmyxListpartsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSmyxListpartsAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSmyxListpartsAPIResponseModel is 药店查询往来单位 成功返回结果
type AlibabaAlihealthDrugKytSmyxListpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_smyx_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSmyxListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSmyxListpartsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSmyxListpartsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSmyxListpartsAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSmyxListpartsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSmyxListpartsAPIResponse
func GetAlibabaAlihealthDrugKytSmyxListpartsAPIResponse() *AlibabaAlihealthDrugKytSmyxListpartsAPIResponse {
	return poolAlibabaAlihealthDrugKytSmyxListpartsAPIResponse.Get().(*AlibabaAlihealthDrugKytSmyxListpartsAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSmyxListpartsAPIResponse 将 AlibabaAlihealthDrugKytSmyxListpartsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSmyxListpartsAPIResponse(v *AlibabaAlihealthDrugKytSmyxListpartsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSmyxListpartsAPIResponse.Put(v)
}
