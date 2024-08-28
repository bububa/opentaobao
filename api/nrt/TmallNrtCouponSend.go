package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtCouponSend 券发放接口
// tmall.nrt.coupon.send
//
// 新零售场景，商家自有渠道发放券
func TmallNrtCouponSend(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtCouponSendAPIRequest, resp *nrt.TmallNrtCouponSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
