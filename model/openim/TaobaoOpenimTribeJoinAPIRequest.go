package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribejoinAPIRequest OPENIM群主动加入 API请求
// taobao.openim.tribe.join
//
// OPENIM群主动加入
type TaobaoopenimtribejoinAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoopenimtribejoinRequest 初始化TaobaoopenimtribejoinAPIRequest对象
func NewTaobaoopenimtribejoinRequest() *TaobaoopenimtribejoinAPIRequest {
	return &TaobaoopenimtribejoinAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribejoinAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.join"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribejoinAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribejoinAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribejoinAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribejoinAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribejoinAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribejoinAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
