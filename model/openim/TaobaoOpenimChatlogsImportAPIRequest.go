package openim

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimChatlogsImportAPIRequest) Reset() {
	r._messages = r._messages[:0]
	r.Params.ToZero()
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

var poolTaobaoOpenimChatlogsImportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimChatlogsImportRequest()
	},
}

// GetTaobaoOpenimChatlogsImportRequest 从 sync.Pool 获取 TaobaoOpenimChatlogsImportAPIRequest
func GetTaobaoOpenimChatlogsImportAPIRequest() *TaobaoOpenimChatlogsImportAPIRequest {
	return poolTaobaoOpenimChatlogsImportAPIRequest.Get().(*TaobaoOpenimChatlogsImportAPIRequest)
}

// ReleaseTaobaoOpenimChatlogsImportAPIRequest 将 TaobaoOpenimChatlogsImportAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimChatlogsImportAPIRequest(v *TaobaoOpenimChatlogsImportAPIRequest) {
	v.Reset()
	poolTaobaoOpenimChatlogsImportAPIRequest.Put(v)
}
