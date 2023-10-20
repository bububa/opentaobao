package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// TaobaoJushitaJmsUserAdd 添加ONS消息同步用户
// taobao.jushita.jms.user.add
//
// 添加ONS消息同步用户
func TaobaoJushitaJmsUserAdd(clt *core.SDKClient, req *jms.TaobaoJushitaJmsUserAddAPIRequest, resp *jms.TaobaoJushitaJmsUserAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
