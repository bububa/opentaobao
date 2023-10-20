package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeSendmsgAPIRequest 发送群消息 API请求
// taobao.openim.tribe.sendmsg
//
// 发送群消息，目前支持发送4种类型的群消息，普通文本，图片，语音，自定义消息
type TaobaoOpenimTribeSendmsgAPIRequest struct {
	model.Params
	// 群消息发送者，只有该群的成员才可以发送群消息
	_user *User
	// 群id
	_tribeId int64
	// 发送群消息
	_msg *TribeMsg
}

// NewTaobaoOpenimTribeSendmsgRequest 初始化TaobaoOpenimTribeSendmsgAPIRequest对象
func NewTaobaoOpenimTribeSendmsgRequest() *TaobaoOpenimTribeSendmsgAPIRequest {
	return &TaobaoOpenimTribeSendmsgAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeSendmsgAPIRequest) Reset() {
	r._user = nil
	r._tribeId = 0
	r._msg = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.sendmsg"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 群消息发送者，只有该群的成员才可以发送群消息
func (r *TaobaoOpenimTribeSendmsgAPIRequest) SetUser(_user *User) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetUser() *User {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeSendmsgAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

// SetMsg is Msg Setter
// 发送群消息
func (r *TaobaoOpenimTribeSendmsgAPIRequest) SetMsg(_msg *TribeMsg) error {
	r._msg = _msg
	r.Set("msg", _msg)
	return nil
}

// GetMsg Msg Getter
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetMsg() *TribeMsg {
	return r._msg
}

var poolTaobaoOpenimTribeSendmsgAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeSendmsgRequest()
	},
}

// GetTaobaoOpenimTribeSendmsgRequest 从 sync.Pool 获取 TaobaoOpenimTribeSendmsgAPIRequest
func GetTaobaoOpenimTribeSendmsgAPIRequest() *TaobaoOpenimTribeSendmsgAPIRequest {
	return poolTaobaoOpenimTribeSendmsgAPIRequest.Get().(*TaobaoOpenimTribeSendmsgAPIRequest)
}

// ReleaseTaobaoOpenimTribeSendmsgAPIRequest 将 TaobaoOpenimTribeSendmsgAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeSendmsgAPIRequest(v *TaobaoOpenimTribeSendmsgAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeSendmsgAPIRequest.Put(v)
}
