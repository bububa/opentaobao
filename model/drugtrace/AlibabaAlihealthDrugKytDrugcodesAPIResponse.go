package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrugcodesAPIResponse 药品是否赋码 API返回值
// alibaba.alihealth.drug.kyt.drugcodes
//
// 药品是否赋码
type AlibabaAlihealthDrugKytDrugcodesAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrugcodesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrugcodesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrugcodesAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrugcodesAPIResponseModel is 药品是否赋码 成功返回结果
type AlibabaAlihealthDrugKytDrugcodesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_drugcodes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrugcodesResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrugcodesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrugcodesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrugcodesAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrugcodesAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrugcodesAPIResponse
func GetAlibabaAlihealthDrugKytDrugcodesAPIResponse() *AlibabaAlihealthDrugKytDrugcodesAPIResponse {
	return poolAlibabaAlihealthDrugKytDrugcodesAPIResponse.Get().(*AlibabaAlihealthDrugKytDrugcodesAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrugcodesAPIResponse 将 AlibabaAlihealthDrugKytDrugcodesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrugcodesAPIResponse(v *AlibabaAlihealthDrugKytDrugcodesAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrugcodesAPIResponse.Put(v)
}
