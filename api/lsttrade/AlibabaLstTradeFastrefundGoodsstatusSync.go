package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// Alibabalsttradefastrefundgoodsstatussync 卖家退款单商品状态同步
// alibaba.lst.trade.fastrefund.goodsstatus.sync
//
// 卖家退款单商品状态同步
func Alibabalsttradefastrefundgoodsstatussync(clt *core.SDKClient, req *lsttrade.AlibabalsttradefastrefundgoodsstatussyncAPIRequest, session string) (*lsttrade.AlibabalsttradefastrefundgoodsstatussyncAPIResponse, error) {
	var resp lsttrade.AlibabalsttradefastrefundgoodsstatussyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
