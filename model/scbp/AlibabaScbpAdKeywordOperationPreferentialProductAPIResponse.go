package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordoperationpreferentialproductAPIResponse 操作优推品 API返回值
// alibaba.scbp.ad.keyword.operation.preferential.product
//
// 操作优推品
type AlibabascbpadkeywordoperationpreferentialproductAPIResponse struct {
	model.CommonResponse
	AlibabascbpadkeywordoperationpreferentialproductAPIResponseModel
}

// AlibabascbpadkeywordoperationpreferentialproductAPIResponseModel is 操作优推品 成功返回结果
type AlibabascbpadkeywordoperationpreferentialproductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_operation_preferential_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功数量
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
