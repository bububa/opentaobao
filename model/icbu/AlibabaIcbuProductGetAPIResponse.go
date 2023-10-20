package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductGetAPIResponse 获得单个商品详情 API返回值
// alibaba.icbu.product.get
//
// 获取商品详情
type AlibabaIcbuProductGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductGetAPIResponseModel).Reset()
}

// AlibabaIcbuProductGetAPIResponseModel is 获得单个商品详情 成功返回结果
type AlibabaIcbuProductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单个商品详情
	Product *AlibabaProductResponse `json:"product,omitempty" xml:"product,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Product = nil
}

var poolAlibabaIcbuProductGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductGetAPIResponse)
	},
}

// GetAlibabaIcbuProductGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductGetAPIResponse
func GetAlibabaIcbuProductGetAPIResponse() *AlibabaIcbuProductGetAPIResponse {
	return poolAlibabaIcbuProductGetAPIResponse.Get().(*AlibabaIcbuProductGetAPIResponse)
}

// ReleaseAlibabaIcbuProductGetAPIResponse 将 AlibabaIcbuProductGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductGetAPIResponse(v *AlibabaIcbuProductGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductGetAPIResponse.Put(v)
}
