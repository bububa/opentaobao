package icbudropshipping

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShippingFreightCalculateAPIResponse 阿里巴巴商品运费计算查询接口 API返回值
// alibaba.shipping.freight.calculate
//
// 阿里巴巴商品运费计算查询接口
type AlibabaShippingFreightCalculateAPIResponse struct {
	model.CommonResponse
	AlibabaShippingFreightCalculateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaShippingFreightCalculateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaShippingFreightCalculateAPIResponseModel).Reset()
}

// AlibabaShippingFreightCalculateAPIResponseModel is 阿里巴巴商品运费计算查询接口 成功返回结果
type AlibabaShippingFreightCalculateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shipping_freight_calculate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// pojo
	Values []Value `json:"values,omitempty" xml:"values>value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaShippingFreightCalculateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Values = m.Values[:0]
}

var poolAlibabaShippingFreightCalculateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaShippingFreightCalculateAPIResponse)
	},
}

// GetAlibabaShippingFreightCalculateAPIResponse 从 sync.Pool 获取 AlibabaShippingFreightCalculateAPIResponse
func GetAlibabaShippingFreightCalculateAPIResponse() *AlibabaShippingFreightCalculateAPIResponse {
	return poolAlibabaShippingFreightCalculateAPIResponse.Get().(*AlibabaShippingFreightCalculateAPIResponse)
}

// ReleaseAlibabaShippingFreightCalculateAPIResponse 将 AlibabaShippingFreightCalculateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaShippingFreightCalculateAPIResponse(v *AlibabaShippingFreightCalculateAPIResponse) {
	v.Reset()
	poolAlibabaShippingFreightCalculateAPIResponse.Put(v)
}
