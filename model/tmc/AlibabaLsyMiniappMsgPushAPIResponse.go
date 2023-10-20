package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsyminiappmsgpushAPIResponse 零售云小程序消息推送 API返回值
// alibaba.lsy.miniapp.msg.push
//
// 零售云小程序消息推送，推送消息至零售云（喵零等）
type AlibabalsyminiappmsgpushAPIResponse struct {
	model.CommonResponse
	AlibabalsyminiappmsgpushAPIResponseModel
}

// AlibabalsyminiappmsgpushAPIResponseModel is 零售云小程序消息推送 成功返回结果
type AlibabalsyminiappmsgpushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_miniapp_msg_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
