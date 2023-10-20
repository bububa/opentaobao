package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltradeshipaddressupdate Openmall订单收货地址修改
// taobao.openmall.trade.shipaddress.update
//
// Openmall订单收货地址修改
func Taobaoopenmalltradeshipaddressupdate(clt *core.SDKClient, req *openmall.TaobaoopenmalltradeshipaddressupdateAPIRequest, session string) (*openmall.TaobaoopenmalltradeshipaddressupdateAPIResponse, error) {
	var resp openmall.TaobaoopenmalltradeshipaddressupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
