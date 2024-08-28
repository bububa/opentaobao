package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvRefundQuery 闲鱼已验货交易订单退款信息查询
// alibaba.idle.isv.refund.query
//
// 闲鱼服务商交易订单退款信息查询-单个退款查询
func AlibabaIdleIsvRefundQuery(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvRefundQueryAPIRequest, resp *idleisv.AlibabaIdleIsvRefundQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
