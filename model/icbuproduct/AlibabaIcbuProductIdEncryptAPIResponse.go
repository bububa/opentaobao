package icbuproduct

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductIdEncryptAPIResponse ICBU国际站商品加密接口 API返回值
// alibaba.icbu.product.id.encrypt
//
// ICBU国际站，对混淆的产品ID加密。
type AlibabaIcbuProductIdEncryptAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductIdEncryptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductIdEncryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductIdEncryptAPIResponseModel).Reset()
}

// AlibabaIcbuProductIdEncryptAPIResponseModel is ICBU国际站商品加密接口 成功返回结果
type AlibabaIcbuProductIdEncryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_id_encrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加密id
	SecretId string `json:"secret_id,omitempty" xml:"secret_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductIdEncryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SecretId = ""
}

var poolAlibabaIcbuProductIdEncryptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductIdEncryptAPIResponse)
	},
}

// GetAlibabaIcbuProductIdEncryptAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductIdEncryptAPIResponse
func GetAlibabaIcbuProductIdEncryptAPIResponse() *AlibabaIcbuProductIdEncryptAPIResponse {
	return poolAlibabaIcbuProductIdEncryptAPIResponse.Get().(*AlibabaIcbuProductIdEncryptAPIResponse)
}

// ReleaseAlibabaIcbuProductIdEncryptAPIResponse 将 AlibabaIcbuProductIdEncryptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductIdEncryptAPIResponse(v *AlibabaIcbuProductIdEncryptAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductIdEncryptAPIResponse.Put(v)
}
