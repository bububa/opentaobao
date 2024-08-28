package lsttrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// AlibabaLstTradeFastrefundGoodsstatusSync 卖家退款单商品状态同步
// alibaba.lst.trade.fastrefund.goodsstatus.sync
//
// 卖家退款单商品状态同步
func AlibabaLstTradeFastrefundGoodsstatusSync(ctx context.Context, clt *core.SDKClient, req *lsttrade.AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest, resp *lsttrade.AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
