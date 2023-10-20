package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrSinglerelationAPIResponse 多融单码关联关系查询 API返回值
// alibaba.alihealth.drug.kyt.dr.singlerelation
//
// 单码关联关系查询
type AlibabaAlihealthDrugKytDrSinglerelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrSinglerelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrSinglerelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrSinglerelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrSinglerelationAPIResponseModel is 多融单码关联关系查询 成功返回结果
type AlibabaAlihealthDrugKytDrSinglerelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_singlerelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrSinglerelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrSinglerelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrSinglerelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrSinglerelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrSinglerelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrSinglerelationAPIResponse
func GetAlibabaAlihealthDrugKytDrSinglerelationAPIResponse() *AlibabaAlihealthDrugKytDrSinglerelationAPIResponse {
	return poolAlibabaAlihealthDrugKytDrSinglerelationAPIResponse.Get().(*AlibabaAlihealthDrugKytDrSinglerelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrSinglerelationAPIResponse 将 AlibabaAlihealthDrugKytDrSinglerelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrSinglerelationAPIResponse(v *AlibabaAlihealthDrugKytDrSinglerelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrSinglerelationAPIResponse.Put(v)
}
