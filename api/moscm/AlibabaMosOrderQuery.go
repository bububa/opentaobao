package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosorderquery 批量查询订单信息
// alibaba.mos.order.query
//
// 查询多笔交易信息
func Alibabamosorderquery(clt *core.SDKClient, req *moscm.AlibabamosorderqueryAPIRequest, session string) (*moscm.AlibabamosorderqueryAPIResponse, error) {
	var resp moscm.AlibabamosorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
