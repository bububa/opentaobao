package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundOrderCancel 取消出库单
// alibaba.tcwms.outbound.order.cancel
//
// 取消出库单
func AlibabaTcwmsOutboundOrderCancel(clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundOrderCancelAPIRequest, resp *perfect.AlibabaTcwmsOutboundOrderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
