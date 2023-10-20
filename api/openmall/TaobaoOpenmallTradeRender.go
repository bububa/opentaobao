package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltraderender 渲染订单价格
// taobao.openmall.trade.render
//
// 请求渲染订单价格
func Taobaoopenmalltraderender(clt *core.SDKClient, req *openmall.TaobaoopenmalltraderenderAPIRequest, session string) (*openmall.TaobaoopenmalltraderenderAPIResponse, error) {
	var resp openmall.TaobaoopenmalltraderenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
