package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJushitaJdpUserAdd 添加数据推送用户
// taobao.jushita.jdp.user.add
//
// 提供给接入数据推送的应用添加数据推送服务的用户
func TaobaoJushitaJdpUserAdd(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJushitaJdpUserAddAPIRequest, resp *jst.TaobaoJushitaJdpUserAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
