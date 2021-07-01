package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJmsUserGetAPIRequest
查询某个用户是否同步消息 API请求
taobao.jushita.jms.user.get

查询某个用户是否同步消息，只支持单个查询 */
type TaobaoJushitaJmsUserGetAPIRequest struct {
	model.Params
	// 需要查询的用户名
	_userNick string
}

// New
