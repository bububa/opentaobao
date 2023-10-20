package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Cainiaorefundrefundactionsget 判断该订单能执行的逆向操作集合列表
// cainiao.refund.refundactions.get
//
// 判断该订单能执行的逆向操作集合列表
func Cainiaorefundrefundactionsget(clt *core.SDKClient, req *trade.CainiaorefundrefundactionsgetAPIRequest, session string) (*trade.CainiaorefundrefundactionsgetAPIResponse, error) {
	var resp trade.CainiaorefundrefundactionsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
