package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymsteamrefundauditAPIResponse 交易猫steam逆向单审核 API返回值
// alibaba.jym.steam.refund.audit
//
// 交易猫steam逆向单审核
type AlibabajymsteamrefundauditAPIResponse struct {
	model.CommonResponse
	AlibabajymsteamrefundauditAPIResponseModel
}

// AlibabajymsteamrefundauditAPIResponseModel is 交易猫steam逆向单审核 成功返回结果
type AlibabajymsteamrefundauditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_steam_refund_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子集错误码，提供与业务细节使用
	SubCodeType string `json:"sub_code_type,omitempty" xml:"sub_code_type,omitempty"`
	// 子集错误描述，提供与业务细节使用
	SubExtraErrMsg string `json:"sub_extra_err_msg,omitempty" xml:"sub_extra_err_msg,omitempty"`
	// 错误码枚举
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误详细描述
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 返回值
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 调用结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
