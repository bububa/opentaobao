package ticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ticket"
)

// AlitripTicketProductQuery 【门票API2.0】门票商品查询接口
// alitrip.ticket.product.query
//
// 门票商品查询接口：返回商家上传的门票商品信息
func AlitripTicketProductQuery(ctx context.Context, clt *core.SDKClient, req *ticket.AlitripTicketProductQueryAPIRequest, resp *ticket.AlitripTicketProductQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
