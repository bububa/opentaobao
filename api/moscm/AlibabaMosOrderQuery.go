package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosOrderQuery 批量查询订单信息
// alibaba.mos.order.query
//
// 查询多笔交易信息
func AlibabaMosOrderQuery(clt *core.SDKClient, req *moscm.AlibabaMosOrderQueryAPIRequest, session string) (*moscm.AlibabaMosOrderQueryAPIResponse, error) {
	var resp moscm.AlibabaMosOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
