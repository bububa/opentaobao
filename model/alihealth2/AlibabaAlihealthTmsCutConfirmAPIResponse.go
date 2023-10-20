package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtmscutconfirmAPIResponse 配拦截失败CP确认结果并回告 API返回值
// alibaba.alihealth.tms.cut.confirm
//
// 配拦截失败CP确认结果并回告
type AlibabaalihealthtmscutconfirmAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtmscutconfirmAPIResponseModel
}

// AlibabaalihealthtmscutconfirmAPIResponseModel is 配拦截失败CP确认结果并回告 成功返回结果
type AlibabaalihealthtmscutconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tms_cut_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihealthtmscutconfirmResult `json:"result,omitempty" xml:"result,omitempty"`
}
