package deliveryvoucher

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherCancelvoucher 作废券
// taobao.game.deliveryvoucher.cancelvoucher
//
// 提货券发券接口：同步券和订单的关联信息
func TaobaoGameDeliveryvoucherCancelvoucher(ctx context.Context, clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherCancelvoucherAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherCancelvoucherAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
