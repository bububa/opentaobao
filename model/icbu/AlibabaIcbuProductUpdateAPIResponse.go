package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductUpdateAPIResponse 修改商品 API返回值
// alibaba.icbu.product.update
//
// 修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
type AlibabaIcbuProductUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductUpdateAPIResponseModel).Reset()
}

// AlibabaIcbuProductUpdateAPIResponseModel is 修改商品 成功返回结果
type AlibabaIcbuProductUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加密后的产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductId = ""
}

var poolAlibabaIcbuProductUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductUpdateAPIResponse)
	},
}

// GetAlibabaIcbuProductUpdateAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductUpdateAPIResponse
func GetAlibabaIcbuProductUpdateAPIResponse() *AlibabaIcbuProductUpdateAPIResponse {
	return poolAlibabaIcbuProductUpdateAPIResponse.Get().(*AlibabaIcbuProductUpdateAPIResponse)
}

// ReleaseAlibabaIcbuProductUpdateAPIResponse 将 AlibabaIcbuProductUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductUpdateAPIResponse(v *AlibabaIcbuProductUpdateAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductUpdateAPIResponse.Put(v)
}
