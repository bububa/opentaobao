package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosOrderListGet 批量查询订单交易
// alibaba.mos.order.list.get
//
// 批量查询交易信息
func AlibabaMosOrderListGet(clt *core.SDKClient, req *moscm.AlibabaMosOrderListGetAPIRequest, session string) (*moscm.AlibabaMosOrderListGetAPIResponse, error) {
	var resp moscm.AlibabaMosOrderListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
