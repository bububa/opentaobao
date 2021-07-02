package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// AlibabaLstTradeFastrefundGoodsstatusSync 卖家退款单商品状态同步
// alibaba.lst.trade.fastrefund.goodsstatus.sync
//
// 卖家退款单商品状态同步
func AlibabaLstTradeFastrefundGoodsstatusSync(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest, session string) (*lsttrade.AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse, error) {
	var resp lsttrade.AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
