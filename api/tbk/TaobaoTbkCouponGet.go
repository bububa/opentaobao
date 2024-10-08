package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkCouponGet 淘宝客-公用-阿里妈妈推广券详情查询
// taobao.tbk.coupon.get
//
// 传入商品ID+券ID(券ID已知情况下)，或者传入me参数，均可查询阿里妈妈推广券详细信息。
func TaobaoTbkCouponGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkCouponGetAPIRequest, resp *tbk.TaobaoTbkCouponGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
