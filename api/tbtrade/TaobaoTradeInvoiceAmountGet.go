package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotradeinvoiceamountget 获取订单应开票金额
// taobao.trade.invoice.amount.get
//
// 订单应开票金额计算
func Taobaotradeinvoiceamountget(clt *core.SDKClient, req *tbtrade.TaobaotradeinvoiceamountgetAPIRequest, session string) (*tbtrade.TaobaotradeinvoiceamountgetAPIResponse, error) {
	var resp tbtrade.TaobaotradeinvoiceamountgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
