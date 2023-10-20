package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// TaobaoJushitaJmsGroupGet 查询ONS分组
// taobao.jushita.jms.group.get
//
// 查询当前appkey在ONS中已有的分组
func TaobaoJushitaJmsGroupGet(clt *core.SDKClient, req *jms.TaobaoJushitaJmsGroupGetAPIRequest, resp *jms.TaobaoJushitaJmsGroupGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
