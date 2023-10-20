package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceOrderGet 天猫国际订单清关信息
// tmall.hk.clearance.order.get
//
// 天猫国际订单清关信息
func TmallHkClearanceOrderGet(clt *core.SDKClient, req *tmallhk.TmallHkClearanceOrderGetAPIRequest, resp *tmallhk.TmallHkClearanceOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
