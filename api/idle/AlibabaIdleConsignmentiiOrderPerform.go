package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleconsignmentiiorderperform 寄卖V2订单履约
// alibaba.idle.consignmentii.order.perform
//
// 寄卖V2订单履约，服务商同步订单信息，驱动交易流转
func Alibabaidleconsignmentiiorderperform(clt *core.SDKClient, req *idle.AlibabaidleconsignmentiiorderperformAPIRequest, session string) (*idle.AlibabaidleconsignmentiiorderperformAPIResponse, error) {
	var resp idle.AlibabaidleconsignmentiiorderperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
