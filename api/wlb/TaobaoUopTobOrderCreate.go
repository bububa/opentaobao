package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoUopTobOrderCreate ToB仓储发货
// taobao.uop.tob.order.create
//
// ToB仓储发货
func TaobaoUopTobOrderCreate(clt *core.SDKClient, req *wlb.TaobaoUopTobOrderCreateAPIRequest, resp *wlb.TaobaoUopTobOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
