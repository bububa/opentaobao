package deliveryvoucher

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherSendvoucher 提货券发券接口
// taobao.game.deliveryvoucher.sendvoucher
//
// 提货券发券接口：同步券和订单的关联信息
func TaobaoGameDeliveryvoucherSendvoucher(ctx context.Context, clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherSendvoucherAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherSendvoucherAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
