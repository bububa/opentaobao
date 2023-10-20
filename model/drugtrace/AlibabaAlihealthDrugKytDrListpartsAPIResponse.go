package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrListpartsAPIResponse 多融查询一个企业的往来单位 API返回值
// alibaba.alihealth.drug.kyt.dr.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugKytDrListpartsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrListpartsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrListpartsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrListpartsAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrListpartsAPIResponseModel is 多融查询一个企业的往来单位 成功返回结果
type AlibabaAlihealthDrugKytDrListpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrListpartsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrListpartsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrListpartsAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrListpartsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrListpartsAPIResponse
func GetAlibabaAlihealthDrugKytDrListpartsAPIResponse() *AlibabaAlihealthDrugKytDrListpartsAPIResponse {
	return poolAlibabaAlihealthDrugKytDrListpartsAPIResponse.Get().(*AlibabaAlihealthDrugKytDrListpartsAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrListpartsAPIResponse 将 AlibabaAlihealthDrugKytDrListpartsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrListpartsAPIResponse(v *AlibabaAlihealthDrugKytDrListpartsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrListpartsAPIResponse.Put(v)
}
