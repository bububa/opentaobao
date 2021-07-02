package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimUsersDeleteAPIRequest 删除用户 API请求
// taobao.openim.users.delete
//
// 批量删除用户
type TaobaoOpenimUsersDeleteAPIRequest struct {
	model.Params
	// 需要删除的用户列表，多个用户用半角逗号分隔，最多一次可以删除100个用户
	_userids []string
}

// NewTaobaoOpenimUsersDeleteRequest 初始化TaobaoOpenimUsersDeleteAPIRequest对象
func NewTaobaoOpenimUsersDeleteRequest() *TaobaoOpenimUsersDeleteAPIRequest {
	return &TaobaoOpenimUsersDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimUsersDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.openim.users.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimUsersDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserids is Userids Setter
// 需要删除的用户列表，多个用户用半角逗号分隔，最多一次可以删除100个用户
func (r *TaobaoOpenimUsersDeleteAPIRequest) SetUserids(_userids []string) error {
	r._userids = _userids
	r.Set("userids", _userids)
	return nil
}

// GetUserids Userids Getter
func (r TaobaoOpenimUsersDeleteAPIRequest) GetUserids() []string {
	return r._userids
}
