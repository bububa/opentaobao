package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqySinglerelationAPIResponse 单码关联关系查询 API返回值
// alibaba.alihealth.drug.kyt.scqy.singlerelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
type AlibabaAlihealthDrugKytScqySinglerelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqySinglerelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqySinglerelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytScqySinglerelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytScqySinglerelationAPIResponseModel is 单码关联关系查询 成功返回结果
type AlibabaAlihealthDrugKytScqySinglerelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_singlerelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytScqySinglerelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqySinglerelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytScqySinglerelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqySinglerelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytScqySinglerelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqySinglerelationAPIResponse
func GetAlibabaAlihealthDrugKytScqySinglerelationAPIResponse() *AlibabaAlihealthDrugKytScqySinglerelationAPIResponse {
	return poolAlibabaAlihealthDrugKytScqySinglerelationAPIResponse.Get().(*AlibabaAlihealthDrugKytScqySinglerelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytScqySinglerelationAPIResponse 将 AlibabaAlihealthDrugKytScqySinglerelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqySinglerelationAPIResponse(v *AlibabaAlihealthDrugKytScqySinglerelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqySinglerelationAPIResponse.Put(v)
}
