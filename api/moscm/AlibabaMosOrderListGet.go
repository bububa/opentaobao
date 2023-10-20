package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosorderlistget 批量查询订单交易
// alibaba.mos.order.list.get
//
// 批量查询交易信息
func Alibabamosorderlistget(clt *core.SDKClient, req *moscm.AlibabamosorderlistgetAPIRequest, session string) (*moscm.AlibabamosorderlistgetAPIResponse, error) {
	var resp moscm.AlibabamosorderlistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
