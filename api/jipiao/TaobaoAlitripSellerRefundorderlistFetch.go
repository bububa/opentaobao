package jipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundorderlistFetch 【机票代理商】——退票订单列表提取
// taobao.alitrip.seller.refundorderlist.fetch
//
// 代理商纬度退票订单列表提取
func TaobaoAlitripSellerRefundorderlistFetch(ctx context.Context, clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundorderlistFetchAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundorderlistFetchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
