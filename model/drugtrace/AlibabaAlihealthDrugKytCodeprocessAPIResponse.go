package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodeprocessAPIResponse 关联关系处理查询 API返回值
// alibaba.alihealth.drug.kyt.codeprocess
//
// 关联关系处理查询
type AlibabaAlihealthDrugKytCodeprocessAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytCodeprocessAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytCodeprocessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytCodeprocessAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytCodeprocessAPIResponseModel is 关联关系处理查询 成功返回结果
type AlibabaAlihealthDrugKytCodeprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_codeprocess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytCodeprocessResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytCodeprocessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytCodeprocessAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytCodeprocessAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytCodeprocessAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytCodeprocessAPIResponse
func GetAlibabaAlihealthDrugKytCodeprocessAPIResponse() *AlibabaAlihealthDrugKytCodeprocessAPIResponse {
	return poolAlibabaAlihealthDrugKytCodeprocessAPIResponse.Get().(*AlibabaAlihealthDrugKytCodeprocessAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytCodeprocessAPIResponse 将 AlibabaAlihealthDrugKytCodeprocessAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytCodeprocessAPIResponse(v *AlibabaAlihealthDrugKytCodeprocessAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytCodeprocessAPIResponse.Put(v)
}
