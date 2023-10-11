package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecycleOrderFulfillSync 同步回收单最终履约方式
// taobao.recycle.order.fulfill.sync
//
// 同步回收单最终履约方式
func TaobaoRecycleOrderFulfillSync(clt *core.SDKClient, req *servicecenter.TaobaoRecycleOrderFulfillSyncAPIRequest, session string) (*servicecenter.TaobaoRecycleOrderFulfillSyncAPIResponse, error) {
	var resp servicecenter.TaobaoRecycleOrderFulfillSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
