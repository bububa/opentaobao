package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribequitAPIRequest OPENIM群成员退出 API请求
// taobao.openim.tribe.quit
//
// OPENIM群成员退出
type TaobaoopenimtribequitAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoopenimtribequitRequest 初始化TaobaoopenimtribequitAPIRequest对象
func NewTaobaoopenimtribequitRequest() *TaobaoopenimtribequitAPIRequest {
	return &TaobaoopenimtribequitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribequitAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.quit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribequitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribequitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribequitAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribequitAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribequitAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribequitAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
