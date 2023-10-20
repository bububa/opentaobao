package icbudropshipping

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOrderFreightCalculateAPIResponse 阿里巴巴下单场景运费方案计算 API返回值
// alibaba.order.freight.calculate
//
// icbu开展 drop shipping 业务，阿里巴巴下单场景运费方案计算
// alibaba Create order scenario freight calculation
type AlibabaOrderFreightCalculateAPIResponse struct {
	model.CommonResponse
	AlibabaOrderFreightCalculateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOrderFreightCalculateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOrderFreightCalculateAPIResponseModel).Reset()
}

// AlibabaOrderFreightCalculateAPIResponseModel is 阿里巴巴下单场景运费方案计算 成功返回结果
type AlibabaOrderFreightCalculateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_order_freight_calculate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Logistics Solution List
	Value []LogisticsSolution `json:"value,omitempty" xml:"value>logistics_solution,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOrderFreightCalculateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = m.Value[:0]
}

var poolAlibabaOrderFreightCalculateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOrderFreightCalculateAPIResponse)
	},
}

// GetAlibabaOrderFreightCalculateAPIResponse 从 sync.Pool 获取 AlibabaOrderFreightCalculateAPIResponse
func GetAlibabaOrderFreightCalculateAPIResponse() *AlibabaOrderFreightCalculateAPIResponse {
	return poolAlibabaOrderFreightCalculateAPIResponse.Get().(*AlibabaOrderFreightCalculateAPIResponse)
}

// ReleaseAlibabaOrderFreightCalculateAPIResponse 将 AlibabaOrderFreightCalculateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOrderFreightCalculateAPIResponse(v *AlibabaOrderFreightCalculateAPIResponse) {
	v.Reset()
	poolAlibabaOrderFreightCalculateAPIResponse.Put(v)
}
