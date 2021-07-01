package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimUsersGetAPIRequest
批量获取用户信息 API请求
taobao.openim.users.get

批量获取用户信息 */
type TaobaoOpenimUsersGetAPIRequest struct {
	model.Params
	// 用户id序列
	_userids []string
}

// New
