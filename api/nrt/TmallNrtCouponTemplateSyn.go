package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtCouponTemplateSyn 喵零券同步
// tmall.nrt.coupon.template.syn
//
// 查询券模版
func TmallNrtCouponTemplateSyn(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtCouponTemplateSynAPIRequest, resp *nrt.TmallNrtCouponTemplateSynAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
