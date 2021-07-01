package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSellercenterRolesGetAPIRequest
获取指定卖家的角色列表 API请求
taobao.sellercenter.roles.get

获取指定卖家的角色列表，只能获取属于登陆者自己的信息。 */
type TaobaoSellercenterRolesGetAPIRequest struct {
	model.Params
	// 卖家昵称(只允许查询自己的信息：当前登陆者)
	_nick string
}

// New
