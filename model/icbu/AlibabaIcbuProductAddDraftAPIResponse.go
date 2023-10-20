package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductAddDraftAPIResponse ICBU商品发布草稿接口 API返回值
// alibaba.icbu.product.add.draft
//
// 发布商品草稿,支持sourcing/一口价商品，支持英文和多种语言原发商品
type AlibabaIcbuProductAddDraftAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductAddDraftAPIResponseModel
}

// AlibabaIcbuProductAddDraftAPIResponseModel is ICBU商品发布草稿接口 成功返回结果
type AlibabaIcbuProductAddDraftAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_add_draft_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 混淆后的产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
