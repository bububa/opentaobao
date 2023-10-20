package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltradeget 查询订单详情
// taobao.openmall.trade.get
//
// 查询订单详情
func Taobaoopenmalltradeget(clt *core.SDKClient, req *openmall.TaobaoopenmalltradegetAPIRequest, session string) (*openmall.TaobaoopenmalltradegetAPIResponse, error) {
	var resp openmall.TaobaoopenmalltradegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
