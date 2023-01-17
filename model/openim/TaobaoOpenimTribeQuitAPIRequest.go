package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeQuitAPIRequest OPENIM群成员退出 API请求
// taobao.openim.tribe.quit
//
// OPENIM群成员退出
type TaobaoOpenimTribeQuitAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoOpenimTribeQuitRequest 初始化TaobaoOpenimTribeQuitAPIRequest对象
func NewTaobaoOpenimTribeQuitRequest() *TaobaoOpenimTribeQuitAPIRequest {
	return &TaobaoOpenimTribeQuitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeQuitAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.quit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeQuitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeQuitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeQuitAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeQuitAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeQuitAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeQuitAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
