package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotagTaguserSave 给用户打上优惠标签
// tmall.promotag.taguser.save
//
// 给用户载体打标
func TmallPromotagTaguserSave(ctx context.Context, clt *core.SDKClient, req *promotion.TmallPromotagTaguserSaveAPIRequest, resp *promotion.TmallPromotagTaguserSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
