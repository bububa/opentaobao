package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmMemberMessageHandleAPIRequest 处理Usc会员消息接口 API请求
// alibaba.westcrm.member.message.handle
//
// 处理Usc会员消息接口
type AlibabaWestcrmMemberMessageHandleAPIRequest struct {
	model.Params
	// 消息类型
	_messageType string
	// 消息内容
	_messageContent string
}

// NewAlibabaWestcrmMemberMessageHandleRequest 初始化AlibabaWestcrmMemberMessageHandleAPIRequest对象
func NewAlibabaWestcrmMemberMessageHandleRequest() *AlibabaWestcrmMemberMessageHandleAPIRequest {
	return &AlibabaWestcrmMemberMessageHandleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmMemberMessageHandleAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.member.message.handle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmMemberMessageHandleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMessageType is MessageType Setter
// 消息类型
func (r *AlibabaWestcrmMemberMessageHandleAPIRequest) SetMessageType(_messageType string) error {
	r._messageType = _messageType
	r.Set("message_type", _messageType)
	return nil
}

// GetMessageType MessageType Getter
func (r AlibabaWestcrmMemberMessageHandleAPIRequest) GetMessageType() string {
	return r._messageType
}

// SetMessageContent is MessageContent Setter
// 消息内容
func (r *AlibabaWestcrmMemberMessageHandleAPIRequest) SetMessageContent(_messageContent string) error {
	r._messageContent = _messageContent
	r.Set("message_content", _messageContent)
	return nil
}

// GetMessageContent MessageContent Getter
func (r AlibabaWestcrmMemberMessageHandleAPIRequest) GetMessageContent() string {
	return r._messageContent
}
