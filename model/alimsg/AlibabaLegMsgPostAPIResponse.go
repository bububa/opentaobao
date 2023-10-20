package alimsg

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegmsgpostAPIResponse 集团法务消息发送 API返回值
// alibaba.leg.msg.post
//
// 消息发送能力
type AlibabalegmsgpostAPIResponse struct {
	model.CommonResponse
	AlibabalegmsgpostAPIResponseModel
}

// AlibabalegmsgpostAPIResponseModel is 集团法务消息发送 成功返回结果
type AlibabalegmsgpostAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_leg_msg_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
