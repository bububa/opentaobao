package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductiddecryptAPIResponse 商品ID解密 API返回值
// alibaba.icbu.product.id.decrypt
//
// 对混淆的产品ID解密
type AlibabaicbuproductiddecryptAPIResponse struct {
	model.CommonResponse
	AlibabaicbuproductiddecryptAPIResponseModel
}

// AlibabaicbuproductiddecryptAPIResponseModel is 商品ID解密 成功返回结果
type AlibabaicbuproductiddecryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_id_decrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
