package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYyListpartsAPIResponse 查询往来单位 API返回值
// alibaba.alihealth.drug.kyt.yy.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugKytYyListpartsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytYyListpartsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytYyListpartsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytYyListpartsAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytYyListpartsAPIResponseModel is 查询往来单位 成功返回结果
type AlibabaAlihealthDrugKytYyListpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yy_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytYyListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytYyListpartsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytYyListpartsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytYyListpartsAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytYyListpartsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytYyListpartsAPIResponse
func GetAlibabaAlihealthDrugKytYyListpartsAPIResponse() *AlibabaAlihealthDrugKytYyListpartsAPIResponse {
	return poolAlibabaAlihealthDrugKytYyListpartsAPIResponse.Get().(*AlibabaAlihealthDrugKytYyListpartsAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytYyListpartsAPIResponse 将 AlibabaAlihealthDrugKytYyListpartsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytYyListpartsAPIResponse(v *AlibabaAlihealthDrugKytYyListpartsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytYyListpartsAPIResponse.Put(v)
}
