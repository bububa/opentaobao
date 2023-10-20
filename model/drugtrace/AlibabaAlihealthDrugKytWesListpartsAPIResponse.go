package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesListpartsAPIResponse 查询往来单位列表 API返回值
// alibaba.alihealth.drug.kyt.wes.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugKytWesListpartsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesListpartsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesListpartsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesListpartsAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesListpartsAPIResponseModel is 查询往来单位列表 成功返回结果
type AlibabaAlihealthDrugKytWesListpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesListpartsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesListpartsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesListpartsAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesListpartsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesListpartsAPIResponse
func GetAlibabaAlihealthDrugKytWesListpartsAPIResponse() *AlibabaAlihealthDrugKytWesListpartsAPIResponse {
	return poolAlibabaAlihealthDrugKytWesListpartsAPIResponse.Get().(*AlibabaAlihealthDrugKytWesListpartsAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesListpartsAPIResponse 将 AlibabaAlihealthDrugKytWesListpartsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesListpartsAPIResponse(v *AlibabaAlihealthDrugKytWesListpartsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesListpartsAPIResponse.Put(v)
}
