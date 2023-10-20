package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaicontentbusinessgetthirdcyclevipstatusAPIResponse 天猫精灵商业化获取三方连续包会员状态 API返回值
// alibaba.ai.content.business.get.third.cycle.vip.status
//
// 天猫精灵商业化获取三方连续包会员状态
type AlibabaaicontentbusinessgetthirdcyclevipstatusAPIResponse struct {
	model.CommonResponse
	AlibabaaicontentbusinessgetthirdcyclevipstatusAPIResponseModel
}

// AlibabaaicontentbusinessgetthirdcyclevipstatusAPIResponseModel is 天猫精灵商业化获取三方连续包会员状态 成功返回结果
type AlibabaaicontentbusinessgetthirdcyclevipstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_content_business_get_third_cycle_vip_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 错误等级
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// 结果码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 三方小会员连续包状态：true-连续包中  false-非连续包
	RetValue bool `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}
