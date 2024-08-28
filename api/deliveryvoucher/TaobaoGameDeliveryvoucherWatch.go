package deliveryvoucher

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherWatch 监控预约数据
// taobao.game.deliveryvoucher.watch
//
// 监控预约数据
func TaobaoGameDeliveryvoucherWatch(ctx context.Context, clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherWatchAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherWatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
