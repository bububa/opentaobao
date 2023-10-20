package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBenefitSendAPIResponse 发奖接口 API返回值
// alibaba.benefit.send
//
// 发奖接口
type AlibabaBenefitSendAPIResponse struct {
	model.CommonResponse
	AlibabaBenefitSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBenefitSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBenefitSendAPIResponseModel).Reset()
}

// AlibabaBenefitSendAPIResponseModel is 发奖接口 成功返回结果
type AlibabaBenefitSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_benefit_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回消息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 接口返回代码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 奖品名称
	PrizeName string `json:"prize_name,omitempty" xml:"prize_name,omitempty"`
	// 权益id
	RightId int64 `json:"right_id,omitempty" xml:"right_id,omitempty"`
	// 是否处理成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBenefitSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
	m.PrizeName = ""
	m.RightId = 0
	m.ResultSuccess = false
}

var poolAlibabaBenefitSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBenefitSendAPIResponse)
	},
}

// GetAlibabaBenefitSendAPIResponse 从 sync.Pool 获取 AlibabaBenefitSendAPIResponse
func GetAlibabaBenefitSendAPIResponse() *AlibabaBenefitSendAPIResponse {
	return poolAlibabaBenefitSendAPIResponse.Get().(*AlibabaBenefitSendAPIResponse)
}

// ReleaseAlibabaBenefitSendAPIResponse 将 AlibabaBenefitSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBenefitSendAPIResponse(v *AlibabaBenefitSendAPIResponse) {
	v.Reset()
	poolAlibabaBenefitSendAPIResponse.Put(v)
}
