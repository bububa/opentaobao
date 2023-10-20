package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse 根据追溯码获取69码 API返回值
// alibaba.alihealth.drug.getbarcode.bytraccode
//
// 根据追溯码获取69码
type AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponseModel is 根据追溯码获取69码 成功返回结果
type AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_getbarcode_bytraccode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse
func GetAlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse() *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse {
	return poolAlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse.Get().(*AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse 将 AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse(v *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse.Put(v)
}
