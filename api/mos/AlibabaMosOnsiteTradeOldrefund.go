package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosonsitetradeoldrefund 线下新退款接口（专为老退款接口调用）
// alibaba.mos.onsite.trade.oldrefund
//
// 线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回
func Alibabamosonsitetradeoldrefund(clt *core.SDKClient, req *mos.AlibabamosonsitetradeoldrefundAPIRequest, session string) (*mos.AlibabamosonsitetradeoldrefundAPIResponse, error) {
	var resp mos.AlibabamosonsitetradeoldrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
