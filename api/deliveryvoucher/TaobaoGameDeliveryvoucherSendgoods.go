package deliveryvoucher

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherSendgoods 提货券发货接口
// taobao.game.deliveryvoucher.sendgoods
//
// 提货券发券接口：同步券和订单的关联信息
func TaobaoGameDeliveryvoucherSendgoods(ctx context.Context, clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherSendgoodsAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherSendgoodsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
