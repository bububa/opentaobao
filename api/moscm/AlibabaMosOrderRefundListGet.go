package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosOrderRefundListGet 批量查询交易退货信息
// alibaba.mos.order.refund.list.get
//
// 批量查询多个退货单的退货明细
func AlibabaMosOrderRefundListGet(clt *core.SDKClient, req *moscm.AlibabaMosOrderRefundListGetAPIRequest, session string) (*moscm.AlibabaMosOrderRefundListGetAPIResponse, error) {
	var resp moscm.AlibabaMosOrderRefundListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
