package jms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// TaobaoJushitaJmsGroupGet 查询ONS分组
// taobao.jushita.jms.group.get
//
// 查询当前appkey在ONS中已有的分组
func TaobaoJushitaJmsGroupGet(ctx context.Context, clt *core.SDKClient, req *jms.TaobaoJushitaJmsGroupGetAPIRequest, resp *jms.TaobaoJushitaJmsGroupGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
