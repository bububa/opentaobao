package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// TaobaoJushitaJmsUserGet 查询某个用户是否同步消息
// taobao.jushita.jms.user.get
//
// 查询某个用户是否同步消息，只支持单个查询
func TaobaoJushitaJmsUserGet(clt *core.SDKClient, req *jms.TaobaoJushitaJmsUserGetAPIRequest, resp *jms.TaobaoJushitaJmsUserGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
