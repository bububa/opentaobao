package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJdsHluserUpdate 订单全链路用户信息修改
// taobao.jds.hluser.update
//
// 订单全链路用户信息修改，比如是否开放买家端展示
func TaobaoJdsHluserUpdate(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJdsHluserUpdateAPIRequest, resp *jst.TaobaoJdsHluserUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
