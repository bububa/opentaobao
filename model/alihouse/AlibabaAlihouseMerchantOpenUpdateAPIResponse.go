package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousemerchantopenupdateAPIResponse 非融合店进件升级成融合店 API返回值
// alibaba.alihouse.merchant.open.update
//
// 非融合店进件升级成融合店
type AlibabaalihousemerchantopenupdateAPIResponse struct {
	model.CommonResponse
	AlibabaalihousemerchantopenupdateAPIResponseModel
}

// AlibabaalihousemerchantopenupdateAPIResponseModel is 非融合店进件升级成融合店 成功返回结果
type AlibabaalihousemerchantopenupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_merchant_open_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousemerchantopenupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
