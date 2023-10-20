package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcmessagesconfirmAPIResponse 确认消费消息的状态 API返回值
// taobao.tmc.messages.confirm
//
// 确认消费消息的状态
type TaobaotmcmessagesconfirmAPIResponse struct {
	model.CommonResponse
	TaobaotmcmessagesconfirmAPIResponseModel
}

// TaobaotmcmessagesconfirmAPIResponseModel is 确认消费消息的状态 成功返回结果
type TaobaotmcmessagesconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_messages_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
