package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvOrderDealrefund 闲鱼无忧购入仓模式服务商退款处理接口
// alibaba.idle.isv.order.dealrefund
//
// 闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款
func AlibabaIdleIsvOrderDealrefund(clt *core.SDKClient, req *idle.AlibabaIdleIsvOrderDealrefundAPIRequest, session string) (*idle.AlibabaIdleIsvOrderDealrefundAPIResponse, error) {
	var resp idle.AlibabaIdleIsvOrderDealrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
