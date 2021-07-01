package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimUsersDeleteAPIRequest
删除用户 API请求
taobao.openim.users.delete

批量删除用户 */
type TaobaoOpenimUsersDeleteAPIRequest struct {
	model.Params
	// 需要删除的用户列表，多个用户用半角逗号分隔，最多一次可以删除100个用户
	_userids []string
}

// New
