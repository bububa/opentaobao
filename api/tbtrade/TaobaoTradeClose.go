package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotradeclose 卖家关闭一笔交易
// taobao.trade.close
//
// 关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
func Taobaotradeclose(clt *core.SDKClient, req *tbtrade.TaobaotradecloseAPIRequest, session string) (*tbtrade.TaobaotradecloseAPIResponse, error) {
	var resp tbtrade.TaobaotradecloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
