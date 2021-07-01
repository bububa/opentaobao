package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJmsUserDeleteAPIRequest
删除ONS消息同步用户 API请求
taobao.jushita.jms.user.delete

删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中 */
type TaobaoJushitaJmsUserDeleteAPIRequest struct {
	model.Params
	// 需要停止同步消息的用户nick
	_userNick string
}

// New
