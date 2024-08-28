package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJushitaJdpUserDelete 删除数据推送用户
// taobao.jushita.jdp.user.delete
//
// 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
func TaobaoJushitaJdpUserDelete(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJushitaJdpUserDeleteAPIRequest, resp *jst.TaobaoJushitaJdpUserDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
