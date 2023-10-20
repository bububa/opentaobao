package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlegoosefishpromotionactivityinfoqueryAPIResponse 闲鱼三方活动参与信息查询 API返回值
// alibaba.idle.goosefish.promotion.activity.info.query
//
// 闲鱼三方活动参与信息查询
type AlibabaidlegoosefishpromotionactivityinfoqueryAPIResponse struct {
	model.CommonResponse
	AlibabaidlegoosefishpromotionactivityinfoqueryAPIResponseModel
}

// AlibabaidlegoosefishpromotionactivityinfoqueryAPIResponseModel is 闲鱼三方活动参与信息查询 成功返回结果
type AlibabaidlegoosefishpromotionactivityinfoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_goosefish_promotion_activity_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ApiErrorCode string `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 错误说明
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// 营销活动参与信息
	Data *PromotionActivityInfoVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}
