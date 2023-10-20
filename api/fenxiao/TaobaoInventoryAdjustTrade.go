package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoinventoryadjusttrade 交易库存调整单
// taobao.inventory.adjust.trade
//
// 商家交易调整库存，淘宝交易、B2B经销等
func Taobaoinventoryadjusttrade(clt *core.SDKClient, req *fenxiao.TaobaoinventoryadjusttradeAPIRequest, session string) (*fenxiao.TaobaoinventoryadjusttradeAPIResponse, error) {
	var resp fenxiao.TaobaoinventoryadjusttradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
