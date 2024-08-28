package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionLimitdiscountDetailGet 限时打折详情查询
// taobao.promotion.limitdiscount.detail.get
//
// 限时打折详情查询。查询出指定限时打折的对应商品记录信息。
func TaobaoPromotionLimitdiscountDetailGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionLimitdiscountDetailGetAPIRequest, resp *promotion.TaobaoPromotionLimitdiscountDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
