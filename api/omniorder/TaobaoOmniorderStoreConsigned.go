package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstoreconsigned Pos端门店发货
// taobao.omniorder.store.consigned
//
// ISV Pos端门店发货，通知星盘
func Taobaoomniorderstoreconsigned(clt *core.SDKClient, req *omniorder.TaobaoomniorderstoreconsignedAPIRequest, session string) (*omniorder.TaobaoomniorderstoreconsignedAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstoreconsignedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
