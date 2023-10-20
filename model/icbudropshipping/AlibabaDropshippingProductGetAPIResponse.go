package icbudropshipping

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDropshippingProductGetAPIResponse 阿里巴巴dropshipping 产品信息获取 API返回值
// alibaba.dropshipping.product.get
//
// 阿里巴巴dropshipping 产品信息获取
type AlibabaDropshippingProductGetAPIResponse struct {
	model.CommonResponse
	AlibabaDropshippingProductGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDropshippingProductGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDropshippingProductGetAPIResponseModel).Reset()
}

// AlibabaDropshippingProductGetAPIResponseModel is 阿里巴巴dropshipping 产品信息获取 成功返回结果
type AlibabaDropshippingProductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dropshipping_product_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// product pojo
	Value []DistributionSaleProduct `json:"value,omitempty" xml:"value>distribution_sale_product,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDropshippingProductGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = m.Value[:0]
}

var poolAlibabaDropshippingProductGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDropshippingProductGetAPIResponse)
	},
}

// GetAlibabaDropshippingProductGetAPIResponse 从 sync.Pool 获取 AlibabaDropshippingProductGetAPIResponse
func GetAlibabaDropshippingProductGetAPIResponse() *AlibabaDropshippingProductGetAPIResponse {
	return poolAlibabaDropshippingProductGetAPIResponse.Get().(*AlibabaDropshippingProductGetAPIResponse)
}

// ReleaseAlibabaDropshippingProductGetAPIResponse 将 AlibabaDropshippingProductGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDropshippingProductGetAPIResponse(v *AlibabaDropshippingProductGetAPIResponse) {
	v.Reset()
	poolAlibabaDropshippingProductGetAPIResponse.Put(v)
}
