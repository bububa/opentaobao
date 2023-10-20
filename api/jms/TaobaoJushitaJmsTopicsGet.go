package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// TaobaoJushitaJmsTopicsGet 根据用户nick获取开通的消息列表
// taobao.jushita.jms.topics.get
//
// 根据用户nick获取开通的消息列表
func TaobaoJushitaJmsTopicsGet(clt *core.SDKClient, req *jms.TaobaoJushitaJmsTopicsGetAPIRequest, resp *jms.TaobaoJushitaJmsTopicsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
