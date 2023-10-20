package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosdeliverysend 发货
// alibaba.mos.delivery.send
//
// 订单发货填写快递单
func Alibabamosdeliverysend(clt *core.SDKClient, req *moscm.AlibabamosdeliverysendAPIRequest, session string) (*moscm.AlibabamosdeliverysendAPIResponse, error) {
	var resp moscm.AlibabamosdeliverysendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
