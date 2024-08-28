package jipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripBuyerGet 敏感信息查询
// taobao.alitrip.buyer.get
//
// 针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。
func TaobaoAlitripBuyerGet(ctx context.Context, clt *core.SDKClient, req *jipiao.TaobaoAlitripBuyerGetAPIRequest, resp *jipiao.TaobaoAlitripBuyerGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
