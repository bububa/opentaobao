package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaorecycleorderfulfillsync 同步回收单最终履约方式
// taobao.recycle.order.fulfill.sync
//
// 同步回收单最终履约方式
func Taobaorecycleorderfulfillsync(clt *core.SDKClient, req *servicecenter.TaobaorecycleorderfulfillsyncAPIRequest, session string) (*servicecenter.TaobaorecycleorderfulfillsyncAPIResponse, error) {
	var resp servicecenter.TaobaorecycleorderfulfillsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
