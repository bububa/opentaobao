package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribegetmembersAPIRequest OPENIM群成员获取 API请求
// taobao.openim.tribe.getmembers
//
// OPENIM群成员获取
type TaobaoopenimtribegetmembersAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoopenimtribegetmembersRequest 初始化TaobaoopenimtribegetmembersAPIRequest对象
func NewTaobaoopenimtribegetmembersRequest() *TaobaoopenimtribegetmembersAPIRequest {
	return &TaobaoopenimtribegetmembersAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribegetmembersAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.getmembers"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribegetmembersAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribegetmembersAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribegetmembersAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribegetmembersAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribegetmembersAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribegetmembersAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
