package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleconsignmentorderget 闲鱼帮卖订单查询
// alibaba.idle.consignment.order.get
//
// 闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息
func Alibabaidleconsignmentorderget(clt *core.SDKClient, req *idle.AlibabaidleconsignmentordergetAPIRequest, session string) (*idle.AlibabaidleconsignmentordergetAPIResponse, error) {
	var resp idle.AlibabaidleconsignmentordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
