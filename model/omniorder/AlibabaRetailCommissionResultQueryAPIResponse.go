package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionResultQueryAPIResponse 分佣结果查询 API返回值
// alibaba.retail.commission.result.query
//
// 查询导购分佣记录
type AlibabaRetailCommissionResultQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailCommissionResultQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailCommissionResultQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailCommissionResultQueryAPIResponseModel).Reset()
}

// AlibabaRetailCommissionResultQueryAPIResponseModel is 分佣结果查询 成功返回结果
type AlibabaRetailCommissionResultQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_commission_result_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailCommissionResultQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailCommissionResultQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailCommissionResultQueryAPIResponse)
	},
}

// GetAlibabaRetailCommissionResultQueryAPIResponse 从 sync.Pool 获取 AlibabaRetailCommissionResultQueryAPIResponse
func GetAlibabaRetailCommissionResultQueryAPIResponse() *AlibabaRetailCommissionResultQueryAPIResponse {
	return poolAlibabaRetailCommissionResultQueryAPIResponse.Get().(*AlibabaRetailCommissionResultQueryAPIResponse)
}

// ReleaseAlibabaRetailCommissionResultQueryAPIResponse 将 AlibabaRetailCommissionResultQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailCommissionResultQueryAPIResponse(v *AlibabaRetailCommissionResultQueryAPIResponse) {
	v.Reset()
	poolAlibabaRetailCommissionResultQueryAPIResponse.Put(v)
}
