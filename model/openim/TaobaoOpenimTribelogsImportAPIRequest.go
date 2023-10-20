package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribelogsimportAPIRequest openim群聊天记录导入 API请求
// taobao.openim.tribelogs.import
//
// openim群聊天记录导入
type TaobaoopenimtribelogsimportAPIRequest struct {
	model.Params
	// 消息列表
	_messages []TribeTextMessage
	// 群号。必须为已存在的群，且群主属于本app
	_tribeId int64
}

// NewTaobaoopenimtribelogsimportRequest 初始化TaobaoopenimtribelogsimportAPIRequest对象
func NewTaobaoopenimtribelogsimportRequest() *TaobaoopenimtribelogsimportAPIRequest {
	return &TaobaoopenimtribelogsimportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribelogsimportAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribelogs.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribelogsimportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribelogsimportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessages is Messages Setter
// 消息列表
func (r *TaobaoopenimtribelogsimportAPIRequest) SetMessages(_messages []TribeTextMessage) error {
	r._messages = _messages
	r.Set("messages", _messages)
	return nil
}

// GetMessages Messages Getter
func (r TaobaoopenimtribelogsimportAPIRequest) GetMessages() []TribeTextMessage {
	return r._messages
}

// SetTribeId is TribeId Setter
// 群号。必须为已存在的群，且群主属于本app
func (r *TaobaoopenimtribelogsimportAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribelogsimportAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
