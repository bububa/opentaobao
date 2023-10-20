package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosOrderQuery 批量查询订单信息
// alibaba.mos.order.query
//
// 查询多笔交易信息
func AlibabaMosOrderQuery(clt *core.SDKClient, req *moscm.AlibabaMosOrderQueryAPIRequest, resp *moscm.AlibabaMosOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
