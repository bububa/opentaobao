package jipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundSearch 【机票代理商】退票申请单列表
// taobao.alitrip.seller.refund.search
//
// 查询退票申请单列表
func TaobaoAlitripSellerRefundSearch(ctx context.Context, clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundSearchAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
