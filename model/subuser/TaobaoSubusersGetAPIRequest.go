package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubusersGetAPIRequest
获取指定账户的子账号简易信息列表 API请求
taobao.subusers.get

获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息） */
type TaobaoSubusersGetAPIRequest struct {
	model.Params
	// 主账号用户名
	_userNick string
}

// New
