package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductIdDecryptAPIResponse 商品ID解密 API返回值
// alibaba.icbu.product.id.decrypt
//
// 对混淆的产品ID解密
type AlibabaIcbuProductIdDecryptAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductIdDecryptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductIdDecryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductIdDecryptAPIResponseModel).Reset()
}

// AlibabaIcbuProductIdDecryptAPIResponseModel is 商品ID解密 成功返回结果
type AlibabaIcbuProductIdDecryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_id_decrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductIdDecryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Id = 0
}

var poolAlibabaIcbuProductIdDecryptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductIdDecryptAPIResponse)
	},
}

// GetAlibabaIcbuProductIdDecryptAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductIdDecryptAPIResponse
func GetAlibabaIcbuProductIdDecryptAPIResponse() *AlibabaIcbuProductIdDecryptAPIResponse {
	return poolAlibabaIcbuProductIdDecryptAPIResponse.Get().(*AlibabaIcbuProductIdDecryptAPIResponse)
}

// ReleaseAlibabaIcbuProductIdDecryptAPIResponse 将 AlibabaIcbuProductIdDecryptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductIdDecryptAPIResponse(v *AlibabaIcbuProductIdDecryptAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductIdDecryptAPIResponse.Put(v)
}
