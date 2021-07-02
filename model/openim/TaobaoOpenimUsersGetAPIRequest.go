package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimUsersGetAPIRequest 批量获取用户信息 API请求
// taobao.openim.users.get
//
// 批量获取用户信息
type TaobaoOpenimUsersGetAPIRequest struct {
	model.Params
	// 用户id序列
	_userids []string
}

// NewTaobaoOpenimUsersGetRequest 初始化TaobaoOpenimUsersGetAPIRequest对象
func NewTaobaoOpenimUsersGetRequest() *TaobaoOpenimUsersGetAPIRequest {
	return &TaobaoOpenimUsersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimUsersGetAPIRequest) GetApiMethodName() string {
	return "taobao.openim.users.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimUsersGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Userids Setter
// 用户id序列
func (r *TaobaoOpenimUsersGetAPIRequest) SetUserids(_userids []string) error {
	r._userids = _userids
	r.Set("userids", _userids)
	return nil
}

// Get Userids Getter
func (r TaobaoOpenimUsersGetAPIRequest) GetUserids() []string {
	return r._userids
}
