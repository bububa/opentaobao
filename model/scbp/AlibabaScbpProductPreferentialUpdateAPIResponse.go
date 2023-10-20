package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpproductpreferentialupdateAPIResponse 设置P4P产品优先推广状态 API返回值
// alibaba.scbp.product.preferential.update
//
// 设置P4P产品优先推广状态
type AlibabascbpproductpreferentialupdateAPIResponse struct {
	model.CommonResponse
	AlibabascbpproductpreferentialupdateAPIResponseModel
}

// AlibabascbpproductpreferentialupdateAPIResponseModel is 设置P4P产品优先推广状态 成功返回结果
type AlibabascbpproductpreferentialupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_product_preferential_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设置成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
