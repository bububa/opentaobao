package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSinglerelationAPIResponse 单码关联关系查询，通过一个码查询这个码下的所有子码 API返回值
// alibaba.alihealth.drug.kyt.singlerelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
type AlibabaAlihealthDrugKytSinglerelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSinglerelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSinglerelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSinglerelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSinglerelationAPIResponseModel is 单码关联关系查询，通过一个码查询这个码下的所有子码 成功返回结果
type AlibabaAlihealthDrugKytSinglerelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_singlerelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytSinglerelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSinglerelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSinglerelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSinglerelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSinglerelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSinglerelationAPIResponse
func GetAlibabaAlihealthDrugKytSinglerelationAPIResponse() *AlibabaAlihealthDrugKytSinglerelationAPIResponse {
	return poolAlibabaAlihealthDrugKytSinglerelationAPIResponse.Get().(*AlibabaAlihealthDrugKytSinglerelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSinglerelationAPIResponse 将 AlibabaAlihealthDrugKytSinglerelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSinglerelationAPIResponse(v *AlibabaAlihealthDrugKytSinglerelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSinglerelationAPIResponse.Put(v)
}
