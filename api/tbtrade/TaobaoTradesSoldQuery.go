package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotradessoldquery 根据收件人信息查询交易单号
// taobao.trades.sold.query
//
// 根据收件人信息查询交易单号。
func Taobaotradessoldquery(clt *core.SDKClient, req *tbtrade.TaobaotradessoldqueryAPIRequest, session string) (*tbtrade.TaobaotradessoldqueryAPIResponse, error) {
	var resp tbtrade.TaobaotradessoldqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
