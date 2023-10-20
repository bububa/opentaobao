package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordlistrelevantproductsAPIResponse 查询和词匹配的推广产品 API返回值
// alibaba.scbp.ad.keyword.list.relevant.products
//
// 查询和词匹配的推广产品
type AlibabascbpadkeywordlistrelevantproductsAPIResponse struct {
	model.CommonResponse
	AlibabascbpadkeywordlistrelevantproductsAPIResponseModel
}

// AlibabascbpadkeywordlistrelevantproductsAPIResponseModel is 查询和词匹配的推广产品 成功返回结果
type AlibabascbpadkeywordlistrelevantproductsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_list_relevant_products_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优推品信息返回
	ResultList []RelevantProductDto `json:"result_list,omitempty" xml:"result_list>relevant_product_dto,omitempty"`
}
