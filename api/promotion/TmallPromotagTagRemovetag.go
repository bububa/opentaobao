package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotagTagRemovetag 删除标签定义
// tmall.promotag.tag.removetag
//
// 用于删除标签定义，但是要确保目前该标签没有人群在使用。
func TmallPromotagTagRemovetag(ctx context.Context, clt *core.SDKClient, req *promotion.TmallPromotagTagRemovetagAPIRequest, resp *promotion.TmallPromotagTagRemovetagAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
