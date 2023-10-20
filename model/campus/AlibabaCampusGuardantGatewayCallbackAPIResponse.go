package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardantgatewaycallbackAPIResponse 人卡关系回调 API返回值
// alibaba.campus.guardant.gateway.callback
//
// 门禁供应商回调平台通知同步结果
type AlibabacampusguardantgatewaycallbackAPIResponse struct {
	model.CommonResponse
	AlibabacampusguardantgatewaycallbackAPIResponseModel
}

// AlibabacampusguardantgatewaycallbackAPIResponseModel is 人卡关系回调 成功返回结果
type AlibabacampusguardantgatewaycallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guardant_gateway_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// PojoResult
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
