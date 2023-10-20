package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlegoosefishuserinfoqueryAPIResponse 闲鱼三方容器用户信息获取 API返回值
// alibaba.idle.goosefish.user.info.query
//
// 闲鱼三方容器用户信息获取
type AlibabaidlegoosefishuserinfoqueryAPIResponse struct {
	model.CommonResponse
	AlibabaidlegoosefishuserinfoqueryAPIResponseModel
}

// AlibabaidlegoosefishuserinfoqueryAPIResponseModel is 闲鱼三方容器用户信息获取 成功返回结果
type AlibabaidlegoosefishuserinfoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_goosefish_user_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ApiErrorCode string `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 错误说明
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// 用户数据
	Data *IdleGooseFishUserInfoVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}
