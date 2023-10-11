package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecyclePredeductSettleSync 同步回收单线下打款明细
// taobao.recycle.prededuct.settle.sync
//
// 同步回收单线下打款明细
func TaobaoRecyclePredeductSettleSync(clt *core.SDKClient, req *servicecenter.TaobaoRecyclePredeductSettleSyncAPIRequest, session string) (*servicecenter.TaobaoRecyclePredeductSettleSyncAPIResponse, error) {
	var resp servicecenter.TaobaoRecyclePredeductSettleSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
