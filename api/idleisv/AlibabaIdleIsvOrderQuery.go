package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvOrderQuery 闲鱼已验货订单查询
// alibaba.idle.isv.order.query
//
// 服务商ISV，根据订单号，查询闲鱼订单信息
func AlibabaIdleIsvOrderQuery(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvOrderQueryAPIRequest, session string) (*idleisv.AlibabaIdleIsvOrderQueryAPIResponse, error) {
	var resp idleisv.AlibabaIdleIsvOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
