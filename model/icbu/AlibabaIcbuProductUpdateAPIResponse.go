package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductupdateAPIResponse 修改商品 API返回值
// alibaba.icbu.product.update
//
// 修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
type AlibabaicbuproductupdateAPIResponse struct {
	model.CommonResponse
	AlibabaicbuproductupdateAPIResponseModel
}

// AlibabaicbuproductupdateAPIResponseModel is 修改商品 成功返回结果
type AlibabaicbuproductupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加密后的产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
