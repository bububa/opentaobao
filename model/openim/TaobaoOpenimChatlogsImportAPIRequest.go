package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimchatlogsimportAPIRequest openim单聊消息导入 API请求
// taobao.openim.chatlogs.import
//
// 提供openim账号的聊天消息导入功能
type TaobaoopenimchatlogsimportAPIRequest struct {
	model.Params
	// 消息序列
	_messages []TextMessage
}

// NewTaobaoopenimchatlogsimportRequest 初始化TaobaoopenimchatlogsimportAPIRequest对象
func NewTaobaoopenimchatlogsimportRequest() *TaobaoopenimchatlogsimportAPIRequest {
	return &TaobaoopenimchatlogsimportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimchatlogsimportAPIRequest) GetApiMethodName() string {
	return "taobao.openim.chatlogs.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimchatlogsimportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimchatlogsimportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessages is Messages Setter
// 消息序列
func (r *TaobaoopenimchatlogsimportAPIRequest) SetMessages(_messages []TextMessage) error {
	r._messages = _messages
	r.Set("messages", _messages)
	return nil
}

// GetMessages Messages Getter
func (r TaobaoopenimchatlogsimportAPIRequest) GetMessages() []TextMessage {
	return r._messages
}
