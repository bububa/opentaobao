package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodetobillAPIResponse 通过追溯码查单据 API返回值
// alibaba.alihealth.drug.kyt.codetobill
//
// 通过追溯码查单据
type AlibabaAlihealthDrugKytCodetobillAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytCodetobillAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytCodetobillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytCodetobillAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytCodetobillAPIResponseModel is 通过追溯码查单据 成功返回结果
type AlibabaAlihealthDrugKytCodetobillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_codetobill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytCodetobillResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytCodetobillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytCodetobillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytCodetobillAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytCodetobillAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytCodetobillAPIResponse
func GetAlibabaAlihealthDrugKytCodetobillAPIResponse() *AlibabaAlihealthDrugKytCodetobillAPIResponse {
	return poolAlibabaAlihealthDrugKytCodetobillAPIResponse.Get().(*AlibabaAlihealthDrugKytCodetobillAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytCodetobillAPIResponse 将 AlibabaAlihealthDrugKytCodetobillAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytCodetobillAPIResponse(v *AlibabaAlihealthDrugKytCodetobillAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytCodetobillAPIResponse.Put(v)
}
