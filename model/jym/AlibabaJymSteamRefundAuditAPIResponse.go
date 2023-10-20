package jym

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymSteamRefundAuditAPIResponse 交易猫steam逆向单审核 API返回值
// alibaba.jym.steam.refund.audit
//
// 交易猫steam逆向单审核
type AlibabaJymSteamRefundAuditAPIResponse struct {
	model.CommonResponse
	AlibabaJymSteamRefundAuditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymSteamRefundAuditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymSteamRefundAuditAPIResponseModel).Reset()
}

// AlibabaJymSteamRefundAuditAPIResponseModel is 交易猫steam逆向单审核 成功返回结果
type AlibabaJymSteamRefundAuditAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaJymSteamRefundAuditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubCodeType = ""
	m.SubExtraErrMsg = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = false
	m.IsSuccess = false
}

var poolAlibabaJymSteamRefundAuditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymSteamRefundAuditAPIResponse)
	},
}

// GetAlibabaJymSteamRefundAuditAPIResponse 从 sync.Pool 获取 AlibabaJymSteamRefundAuditAPIResponse
func GetAlibabaJymSteamRefundAuditAPIResponse() *AlibabaJymSteamRefundAuditAPIResponse {
	return poolAlibabaJymSteamRefundAuditAPIResponse.Get().(*AlibabaJymSteamRefundAuditAPIResponse)
}

// ReleaseAlibabaJymSteamRefundAuditAPIResponse 将 AlibabaJymSteamRefundAuditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymSteamRefundAuditAPIResponse(v *AlibabaJymSteamRefundAuditAPIResponse) {
	v.Reset()
	poolAlibabaJymSteamRefundAuditAPIResponse.Put(v)
}
