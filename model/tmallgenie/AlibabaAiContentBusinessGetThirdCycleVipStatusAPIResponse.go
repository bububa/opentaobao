package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse 天猫精灵商业化获取三方连续包会员状态 API返回值
// alibaba.ai.content.business.get.third.cycle.vip.status
//
// 天猫精灵商业化获取三方连续包会员状态
type AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse struct {
	model.CommonResponse
	AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponseModel).Reset()
}

// AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponseModel is 天猫精灵商业化获取三方连续包会员状态 成功返回结果
type AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetMsg = ""
	m.ErrorLevel = ""
	m.RetCode = 0
	m.RetValue = false
}

var poolAlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse)
	},
}

// GetAlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse 从 sync.Pool 获取 AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse
func GetAlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse() *AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse {
	return poolAlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse.Get().(*AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse)
}

// ReleaseAlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse 将 AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse(v *AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse) {
	v.Reset()
	poolAlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse.Put(v)
}
