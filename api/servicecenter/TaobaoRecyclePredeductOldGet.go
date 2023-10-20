package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecyclePredeductOldGet 查询回收单前置抵扣详情
// taobao.recycle.prededuct.old.get
//
// 查询回收单前置抵扣详情
func TaobaoRecyclePredeductOldGet(clt *core.SDKClient, req *servicecenter.TaobaoRecyclePredeductOldGetAPIRequest, resp *servicecenter.TaobaoRecyclePredeductOldGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
