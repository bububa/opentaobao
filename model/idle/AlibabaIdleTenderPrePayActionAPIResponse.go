package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderprepayactionAPIResponse 服务商预付款完成接口 API返回值
// alibaba.idle.tender.pre.pay.action
//
// 服务商预付款完成接口
type AlibabaidletenderprepayactionAPIResponse struct {
	model.CommonResponse
	AlibabaidletenderprepayactionAPIResponseModel
}

// AlibabaidletenderprepayactionAPIResponseModel is 服务商预付款完成接口 成功返回结果
type AlibabaidletenderprepayactionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_pre_pay_action_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
