package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// TaobaoJushitaJmsUserDelete 删除ONS消息同步用户
// taobao.jushita.jms.user.delete
//
// 删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
func TaobaoJushitaJmsUserDelete(clt *core.SDKClient, req *jms.TaobaoJushitaJmsUserDeleteAPIRequest, resp *jms.TaobaoJushitaJmsUserDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
