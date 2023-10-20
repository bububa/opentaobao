package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceGet 天猫国际-清关材料查询
// tmall.hk.clearance.get
//
// 提供订单收货人身份信息查询功能。
func TmallHkClearanceGet(clt *core.SDKClient, req *tmallhk.TmallHkClearanceGetAPIRequest, resp *tmallhk.TmallHkClearanceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
