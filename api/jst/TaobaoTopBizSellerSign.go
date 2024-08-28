package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoTopBizSellerSign 淘宝订单履约-商家erp签约
// taobao.top.biz.seller.sign
//
// 淘宝订单履约-商家erp签约，包含各场景的签约
func TaobaoTopBizSellerSign(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoTopBizSellerSignAPIRequest, resp *jst.TaobaoTopBizSellerSignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
