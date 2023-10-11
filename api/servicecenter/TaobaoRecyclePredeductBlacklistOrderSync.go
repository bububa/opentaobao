package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecyclePredeductBlacklistOrderSync 同步服务商黑名单
// taobao.recycle.prededuct.blacklist.order.sync
//
// 同步服务商黑名单
func TaobaoRecyclePredeductBlacklistOrderSync(clt *core.SDKClient, req *servicecenter.TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest, session string) (*servicecenter.TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse, error) {
	var resp servicecenter.TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
