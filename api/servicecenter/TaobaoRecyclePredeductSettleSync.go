package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaorecyclepredeductsettlesync 同步回收单线下打款明细
// taobao.recycle.prededuct.settle.sync
//
// 同步回收单线下打款明细
func Taobaorecyclepredeductsettlesync(clt *core.SDKClient, req *servicecenter.TaobaorecyclepredeductsettlesyncAPIRequest, session string) (*servicecenter.TaobaorecyclepredeductsettlesyncAPIResponse, error) {
	var resp servicecenter.TaobaorecyclepredeductsettlesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
