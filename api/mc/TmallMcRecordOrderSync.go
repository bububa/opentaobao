package mc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// TmallMcRecordOrderSync 订单信息同步
// tmall.mc.record.order.sync
//
// 订单信息同步(零售云接口)
func TmallMcRecordOrderSync(clt *core.SDKClient, req *mc.TmallMcRecordOrderSyncAPIRequest, resp *mc.TmallMcRecordOrderSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
