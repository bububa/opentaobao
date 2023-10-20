package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosonsitetradeisnewpayorder 是否为新支付订单
// alibaba.mos.onsite.trade.isnewpayorder
//
// 退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单
func Alibabamosonsitetradeisnewpayorder(clt *core.SDKClient, req *mos.AlibabamosonsitetradeisnewpayorderAPIRequest, session string) (*mos.AlibabamosonsitetradeisnewpayorderAPIResponse, error) {
	var resp mos.AlibabamosonsitetradeisnewpayorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
