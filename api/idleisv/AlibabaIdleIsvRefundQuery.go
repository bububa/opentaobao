package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvRefundQuery 闲鱼已验货交易订单退款信息查询
// alibaba.idle.isv.refund.query
//
// 闲鱼服务商交易订单退款信息查询-单个退款查询
func AlibabaIdleIsvRefundQuery(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvRefundQueryAPIRequest, session string) (*idleisv.AlibabaIdleIsvRefundQueryAPIResponse, error) {
	var resp idleisv.AlibabaIdleIsvRefundQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
