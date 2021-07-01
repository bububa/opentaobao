package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJmsTopicsGetAPIRequest
根据用户nick获取开通的消息列表 API请求
taobao.jushita.jms.topics.get

根据用户nick获取开通的消息列表 */
type TaobaoJushitaJmsTopicsGetAPIRequest struct {
	model.Params
	// 卖家nick
	_nick string
}

// New
