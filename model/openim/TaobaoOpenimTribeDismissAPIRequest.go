package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribedismissAPIRequest OPENIM群解散 API请求
// taobao.openim.tribe.dismiss
//
// OPENIM群解散
type TaobaoopenimtribedismissAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoopenimtribedismissRequest 初始化TaobaoopenimtribedismissAPIRequest对象
func NewTaobaoopenimtribedismissRequest() *TaobaoopenimtribedismissAPIRequest {
	return &TaobaoopenimtribedismissAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribedismissAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.dismiss"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribedismissAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribedismissAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribedismissAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribedismissAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribedismissAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribedismissAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
