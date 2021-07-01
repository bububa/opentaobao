package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimUsersUpdateAPIRequest
批量更新用户信息 API请求
taobao.openim.users.update

批量更新用户信息 */
type TaobaoOpenimUsersUpdateAPIRequest struct {
	model.Params
	// 用户信息列表
	_userinfos []Userinfos
}

// New
