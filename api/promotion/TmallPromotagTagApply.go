package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotagTagApply 优惠标签申请
// tmall.promotag.tag.apply
//
// 创建优惠标签
func TmallPromotagTagApply(ctx context.Context, clt *core.SDKClient, req *promotion.TmallPromotagTagApplyAPIRequest, resp *promotion.TmallPromotagTagApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
