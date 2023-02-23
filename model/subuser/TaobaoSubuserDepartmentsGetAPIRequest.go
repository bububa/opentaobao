package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubuserDepartmentsGetAPIRequest 获取指定账户的所有部门列表 API请求
// taobao.subuser.departments.get
//
// 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
type TaobaoSubuserDepartmentsGetAPIRequest struct {
	model.Params
	// 主账号用户名
	_userNick string
}

// NewTaobaoSubuserDepartmentsGetRequest 初始化TaobaoSubuserDepartmentsGetAPIRequest对象
func NewTaobaoSubuserDepartmentsGetRequest() *TaobaoSubuserDepartmentsGetAPIRequest {
	return &TaobaoSubuserDepartmentsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubuserDepartmentsGetAPIRequest) GetApiMethodName() string {
	return "taobao.subuser.departments.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubuserDepartmentsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubuserDepartmentsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 主账号用户名
func (r *TaobaoSubuserDepartmentsGetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoSubuserDepartmentsGetAPIRequest) GetUserNick() string {
	return r._userNick
}
