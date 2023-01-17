package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimChatlogsImportAPIRequest openim单聊消息导入 API请求
// taobao.openim.chatlogs.import
//
// 提供openim账号的聊天消息导入功能
type TaobaoOpenimChatlogsImportAPIRequest struct {
	model.Params
	// 消息序列
	_messages []TextMessage
}

// NewTaobaoOpenimChatlogsImportRequest 初始化TaobaoOpenimChatlogsImportAPIRequest对象
func NewTaobaoOpenimChatlogsImportRequest() *TaobaoOpenimChatlogsImportAPIRequest {
	return &TaobaoOpenimChatlogsImportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimChatlogsImportAPIRequest) GetApiMethodName() string {
	return "taobao.openim.chatlogs.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimChatlogsImportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimChatlogsImportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessages is Messages Setter
// 消息序列
func (r *TaobaoOpenimChatlogsImportAPIRequest) SetMessages(_messages []TextMessage) error {
	r._messages = _messages
	r.Set("messages", _messages)
	return nil
}

// GetMessages Messages Getter
func (r TaobaoOpenimChatlogsImportAPIRequest) GetMessages() []TextMessage {
	return r._messages
}
