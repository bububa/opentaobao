package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribesendmsgAPIRequest 发送群消息 API请求
// taobao.openim.tribe.sendmsg
//
// 发送群消息，目前支持发送4种类型的群消息，普通文本，图片，语音，自定义消息
type TaobaoopenimtribesendmsgAPIRequest struct {
	model.Params
	// 群消息发送者，只有该群的成员才可以发送群消息
	_user *User
	// 群id
	_tribeId int64
	// 发送群消息
	_msg *TribeMsg
}

// NewTaobaoopenimtribesendmsgRequest 初始化TaobaoopenimtribesendmsgAPIRequest对象
func NewTaobaoopenimtribesendmsgRequest() *TaobaoopenimtribesendmsgAPIRequest {
	return &TaobaoopenimtribesendmsgAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribesendmsgAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.sendmsg"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribesendmsgAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribesendmsgAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 群消息发送者，只有该群的成员才可以发送群消息
func (r *TaobaoopenimtribesendmsgAPIRequest) SetUser(_user *User) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribesendmsgAPIRequest) GetUser() *User {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribesendmsgAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribesendmsgAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

// SetMsg is Msg Setter
// 发送群消息
func (r *TaobaoopenimtribesendmsgAPIRequest) SetMsg(_msg *TribeMsg) error {
	r._msg = _msg
	r.Set("msg", _msg)
	return nil
}

// GetMsg Msg Getter
func (r TaobaoopenimtribesendmsgAPIRequest) GetMsg() *TribeMsg {
	return r._msg
}
