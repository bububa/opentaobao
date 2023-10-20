package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribelogsImportAPIRequest openim群聊天记录导入 API请求
// taobao.openim.tribelogs.import
//
// openim群聊天记录导入
type TaobaoOpenimTribelogsImportAPIRequest struct {
	model.Params
	// 消息列表
	_messages []TribeTextMessage
	// 群号。必须为已存在的群，且群主属于本app
	_tribeId int64
}

// NewTaobaoOpenimTribelogsImportRequest 初始化TaobaoOpenimTribelogsImportAPIRequest对象
func NewTaobaoOpenimTribelogsImportRequest() *TaobaoOpenimTribelogsImportAPIRequest {
	return &TaobaoOpenimTribelogsImportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribelogsImportAPIRequest) Reset() {
	r._messages = r._messages[:0]
	r._tribeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribelogsImportAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribelogs.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribelogsImportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribelogsImportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessages is Messages Setter
// 消息列表
func (r *TaobaoOpenimTribelogsImportAPIRequest) SetMessages(_messages []TribeTextMessage) error {
	r._messages = _messages
	r.Set("messages", _messages)
	return nil
}

// GetMessages Messages Getter
func (r TaobaoOpenimTribelogsImportAPIRequest) GetMessages() []TribeTextMessage {
	return r._messages
}

// SetTribeId is TribeId Setter
// 群号。必须为已存在的群，且群主属于本app
func (r *TaobaoOpenimTribelogsImportAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribelogsImportAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

var poolTaobaoOpenimTribelogsImportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribelogsImportRequest()
	},
}

// GetTaobaoOpenimTribelogsImportRequest 从 sync.Pool 获取 TaobaoOpenimTribelogsImportAPIRequest
func GetTaobaoOpenimTribelogsImportAPIRequest() *TaobaoOpenimTribelogsImportAPIRequest {
	return poolTaobaoOpenimTribelogsImportAPIRequest.Get().(*TaobaoOpenimTribelogsImportAPIRequest)
}

// ReleaseTaobaoOpenimTribelogsImportAPIRequest 将 TaobaoOpenimTribelogsImportAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribelogsImportAPIRequest(v *TaobaoOpenimTribelogsImportAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribelogsImportAPIRequest.Put(v)
}
