package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeListCodeGovAPIResponse 政府码查询接口 API返回值
// alibaba.alihealth.drug.code.list.code.gov
//
// 政府码查询接口
type AlibabaAlihealthDrugCodeListCodeGovAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeListCodeGovAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeListCodeGovAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeListCodeGovAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeListCodeGovAPIResponseModel is 政府码查询接口 成功返回结果
type AlibabaAlihealthDrugCodeListCodeGovAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_list_code_gov_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeListCodeGovResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeListCodeGovAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeListCodeGovAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeListCodeGovAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeListCodeGovAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeListCodeGovAPIResponse
func GetAlibabaAlihealthDrugCodeListCodeGovAPIResponse() *AlibabaAlihealthDrugCodeListCodeGovAPIResponse {
	return poolAlibabaAlihealthDrugCodeListCodeGovAPIResponse.Get().(*AlibabaAlihealthDrugCodeListCodeGovAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeListCodeGovAPIResponse 将 AlibabaAlihealthDrugCodeListCodeGovAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeListCodeGovAPIResponse(v *AlibabaAlihealthDrugCodeListCodeGovAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeListCodeGovAPIResponse.Put(v)
}
