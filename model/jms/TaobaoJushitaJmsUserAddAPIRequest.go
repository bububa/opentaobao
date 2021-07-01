package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJmsUserAddAPIRequest
添加ONS消息同步用户 API请求
taobao.jushita.jms.user.add

添加ONS消息同步用户 */
type TaobaoJushitaJmsUserAddAPIRequest struct {
	model.Params
	// topic列表,不填则继承appkey所订阅的topic
	_topicNames []string
}

// New
