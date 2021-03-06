package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoUopTobOrderCreate ToB仓储发货
// taobao.uop.tob.order.create
//
// ToB仓储发货
func TaobaoUopTobOrderCreate(clt *core.SDKClient, req *wlb.TaobaoUopTobOrderCreateAPIRequest, session string) (*wlb.TaobaoUopTobOrderCreateAPIResponse, error) {
	var resp wlb.TaobaoUopTobOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
