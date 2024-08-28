package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenOrderBackflow 订单回流接口
// alibaba.alsc.crm.open.order.backflow
//
// 回流isv订单接口
func AlibabaAlscCrmOpenOrderBackflow(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenOrderBackflowAPIRequest, resp *alsc.AlibabaAlscCrmOpenOrderBackflowAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
