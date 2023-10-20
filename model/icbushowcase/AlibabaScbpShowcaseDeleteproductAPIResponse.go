package icbushowcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpshowcasedeleteproductAPIResponse 批量删除橱窗商品 API返回值
// alibaba.scbp.showcase.deleteproduct
//
// 批量删除橱窗商品
type AlibabascbpshowcasedeleteproductAPIResponse struct {
	model.CommonResponse
	AlibabascbpshowcasedeleteproductAPIResponseModel
}

// AlibabascbpshowcasedeleteproductAPIResponseModel is 批量删除橱窗商品 成功返回结果
type AlibabascbpshowcasedeleteproductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_deleteproduct_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
