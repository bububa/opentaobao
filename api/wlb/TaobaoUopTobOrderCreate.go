package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaouoptobordercreate ToB仓储发货
// taobao.uop.tob.order.create
//
// ToB仓储发货
func Taobaouoptobordercreate(clt *core.SDKClient, req *wlb.TaobaouoptobordercreateAPIRequest, session string) (*wlb.TaobaouoptobordercreateAPIResponse, error) {
	var resp wlb.TaobaouoptobordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
