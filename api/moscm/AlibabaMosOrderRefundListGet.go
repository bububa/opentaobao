package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosorderrefundlistget 批量查询交易退货信息
// alibaba.mos.order.refund.list.get
//
// 批量查询多个退货单的退货明细
func Alibabamosorderrefundlistget(clt *core.SDKClient, req *moscm.AlibabamosorderrefundlistgetAPIRequest, session string) (*moscm.AlibabamosorderrefundlistgetAPIResponse, error) {
	var resp moscm.AlibabamosorderrefundlistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
