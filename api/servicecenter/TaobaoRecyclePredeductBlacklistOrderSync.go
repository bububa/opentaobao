package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecyclePredeductBlacklistOrderSync 同步服务商黑名单
// taobao.recycle.prededuct.blacklist.order.sync
//
// 同步服务商黑名单
func TaobaoRecyclePredeductBlacklistOrderSync(clt *core.SDKClient, req *servicecenter.TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest, resp *servicecenter.TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
