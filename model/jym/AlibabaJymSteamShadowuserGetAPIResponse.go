package jym

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymSteamShadowuserGetAPIResponse 获取影子标识 API返回值
// alibaba.jym.steam.shadowuser.get
//
// 交易猫Steam类目获取影子ID
type AlibabaJymSteamShadowuserGetAPIResponse struct {
	model.CommonResponse
	AlibabaJymSteamShadowuserGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymSteamShadowuserGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymSteamShadowuserGetAPIResponseModel).Reset()
}

// AlibabaJymSteamShadowuserGetAPIResponseModel is 获取影子标识 成功返回结果
type AlibabaJymSteamShadowuserGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_steam_shadowuser_get_response"`
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
	// 影子DTO
	Result *SteamShadowDto `json:"result,omitempty" xml:"result,omitempty"`
	// 调用结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymSteamShadowuserGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubCodeType = ""
	m.SubExtraErrMsg = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.IsSuccess = false
}

var poolAlibabaJymSteamShadowuserGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymSteamShadowuserGetAPIResponse)
	},
}

// GetAlibabaJymSteamShadowuserGetAPIResponse 从 sync.Pool 获取 AlibabaJymSteamShadowuserGetAPIResponse
func GetAlibabaJymSteamShadowuserGetAPIResponse() *AlibabaJymSteamShadowuserGetAPIResponse {
	return poolAlibabaJymSteamShadowuserGetAPIResponse.Get().(*AlibabaJymSteamShadowuserGetAPIResponse)
}

// ReleaseAlibabaJymSteamShadowuserGetAPIResponse 将 AlibabaJymSteamShadowuserGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymSteamShadowuserGetAPIResponse(v *AlibabaJymSteamShadowuserGetAPIResponse) {
	v.Reset()
	poolAlibabaJymSteamShadowuserGetAPIResponse.Put(v)
}
