package jipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundRefusereturn 【机票代理商】拒绝退票
// taobao.alitrip.seller.refund.refusereturn
//
// 拒绝退票
func TaobaoAlitripSellerRefundRefusereturn(ctx context.Context, clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundRefusereturnAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundRefusereturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
