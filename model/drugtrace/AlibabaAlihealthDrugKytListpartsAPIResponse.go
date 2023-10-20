package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytListpartsAPIResponse 查询往来单位列表 API返回值
// alibaba.alihealth.drug.kyt.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugKytListpartsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytListpartsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytListpartsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytListpartsAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytListpartsAPIResponseModel is 查询往来单位列表 成功返回结果
type AlibabaAlihealthDrugKytListpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytListpartsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytListpartsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytListpartsAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytListpartsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytListpartsAPIResponse
func GetAlibabaAlihealthDrugKytListpartsAPIResponse() *AlibabaAlihealthDrugKytListpartsAPIResponse {
	return poolAlibabaAlihealthDrugKytListpartsAPIResponse.Get().(*AlibabaAlihealthDrugKytListpartsAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytListpartsAPIResponse 将 AlibabaAlihealthDrugKytListpartsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytListpartsAPIResponse(v *AlibabaAlihealthDrugKytListpartsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytListpartsAPIResponse.Put(v)
}
