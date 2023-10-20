package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionOrderQueryAPIResponse 分销订单查询 API返回值
// alibaba.retail.commission.order.query
//
// 查询商家的分销订单
type AlibabaRetailCommissionOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailCommissionOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailCommissionOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailCommissionOrderQueryAPIResponseModel).Reset()
}

// AlibabaRetailCommissionOrderQueryAPIResponseModel is 分销订单查询 成功返回结果
type AlibabaRetailCommissionOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_commission_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailCommissionOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailCommissionOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailCommissionOrderQueryAPIResponse)
	},
}

// GetAlibabaRetailCommissionOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaRetailCommissionOrderQueryAPIResponse
func GetAlibabaRetailCommissionOrderQueryAPIResponse() *AlibabaRetailCommissionOrderQueryAPIResponse {
	return poolAlibabaRetailCommissionOrderQueryAPIResponse.Get().(*AlibabaRetailCommissionOrderQueryAPIResponse)
}

// ReleaseAlibabaRetailCommissionOrderQueryAPIResponse 将 AlibabaRetailCommissionOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailCommissionOrderQueryAPIResponse(v *AlibabaRetailCommissionOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaRetailCommissionOrderQueryAPIResponse.Put(v)
}
