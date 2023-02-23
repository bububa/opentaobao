package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseMessageWorkorderPushAPIRequest 工单消息推送 API请求
// alibaba.alihouse.message.workorder.push
//
// 工单消息推送
type AlibabaAlihouseMessageWorkorderPushAPIRequest struct {
	model.Params
	// 入参
	_messageInfoDto *MessageInfoDto
}

// NewAlibabaAlihouseMessageWorkorderPushRequest 初始化AlibabaAlihouseMessageWorkorderPushAPIRequest对象
func NewAlibabaAlihouseMessageWorkorderPushRequest() *AlibabaAlihouseMessageWorkorderPushAPIRequest {
	return &AlibabaAlihouseMessageWorkorderPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseMessageWorkorderPushAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.message.workorder.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseMessageWorkorderPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseMessageWorkorderPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessageInfoDto is MessageInfoDto Setter
// 入参
func (r *AlibabaAlihouseMessageWorkorderPushAPIRequest) SetMessageInfoDto(_messageInfoDto *MessageInfoDto) error {
	r._messageInfoDto = _messageInfoDto
	r.Set("message_info_dto", _messageInfoDto)
	return nil
}

// GetMessageInfoDto MessageInfoDto Getter
func (r AlibabaAlihouseMessageWorkorderPushAPIRequest) GetMessageInfoDto() *MessageInfoDto {
	return r._messageInfoDto
}
