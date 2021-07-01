package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductAddAPIResponse
发布产品 API返回值
alibaba.icbu.product.add

发布商品,支持sourcing/一口价商品，支持英文和多种语言原发商品 */
type AlibabaIcbuProductAddAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductAddAPIResponseModel
}

// AlibabaIcbuProductAddAPIResponseModel is 发布产品 成功返回结果
type AlibabaIcbuProductAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 混淆后的产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
