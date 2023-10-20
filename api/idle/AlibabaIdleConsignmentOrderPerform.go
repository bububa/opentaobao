package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleconsignmentorderperform 帮卖订单履约
// alibaba.idle.consignment.order.perform
//
// 帮卖订单履约，回收商同步订单信息，驱动交易流转
func Alibabaidleconsignmentorderperform(clt *core.SDKClient, req *idle.AlibabaidleconsignmentorderperformAPIRequest, session string) (*idle.AlibabaidleconsignmentorderperformAPIResponse, error) {
	var resp idle.AlibabaidleconsignmentorderperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
