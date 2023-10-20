package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductUpdateFieldAPIResponse 商品按字段更新 API返回值
// alibaba.icbu.product.update.field
//
// 按字段修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
type AlibabaIcbuProductUpdateFieldAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductUpdateFieldAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductUpdateFieldAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductUpdateFieldAPIResponseModel).Reset()
}

// AlibabaIcbuProductUpdateFieldAPIResponseModel is 商品按字段更新 成功返回结果
type AlibabaIcbuProductUpdateFieldAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_update_field_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加密后的产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductUpdateFieldAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductId = ""
}

var poolAlibabaIcbuProductUpdateFieldAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductUpdateFieldAPIResponse)
	},
}

// GetAlibabaIcbuProductUpdateFieldAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductUpdateFieldAPIResponse
func GetAlibabaIcbuProductUpdateFieldAPIResponse() *AlibabaIcbuProductUpdateFieldAPIResponse {
	return poolAlibabaIcbuProductUpdateFieldAPIResponse.Get().(*AlibabaIcbuProductUpdateFieldAPIResponse)
}

// ReleaseAlibabaIcbuProductUpdateFieldAPIResponse 将 AlibabaIcbuProductUpdateFieldAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductUpdateFieldAPIResponse(v *AlibabaIcbuProductUpdateFieldAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductUpdateFieldAPIResponse.Put(v)
}
