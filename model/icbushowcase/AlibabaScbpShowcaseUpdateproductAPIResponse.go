package icbushowcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpshowcaseupdateproductAPIResponse 替换橱窗商品 API返回值
// alibaba.scbp.showcase.updateproduct
//
// 替换橱窗商品
type AlibabascbpshowcaseupdateproductAPIResponse struct {
	model.CommonResponse
	AlibabascbpshowcaseupdateproductAPIResponseModel
}

// AlibabascbpshowcaseupdateproductAPIResponseModel is 替换橱窗商品 成功返回结果
type AlibabascbpshowcaseupdateproductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_updateproduct_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否修改成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
