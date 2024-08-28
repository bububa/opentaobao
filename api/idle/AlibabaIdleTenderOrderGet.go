package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderOrderGet 暗拍读取订单
// alibaba.idle.tender.order.get
//
// 查询省心卖暗拍项目订单
func AlibabaIdleTenderOrderGet(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTenderOrderGetAPIRequest, resp *idle.AlibabaIdleTenderOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
