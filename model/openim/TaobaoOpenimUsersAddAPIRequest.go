package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimUsersAddAPIRequest
添加用户 API请求
taobao.openim.users.add

导入用户 */
type TaobaoOpenimUsersAddAPIRequest struct {
	model.Params
	// 用户信息列表
	_userinfos []Userinfos
}

// New
