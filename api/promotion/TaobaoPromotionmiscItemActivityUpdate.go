package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscItemActivityUpdate 修改无条件单品优惠活动
// taobao.promotionmisc.item.activity.update
//
// 修改无条件单品优惠活动。&lt;br/&gt;1、该接口只修改活动基本信息和打折信息，如需要增加、删除参与该活动的商品，请调用taobao.promotionmisc.activity.range.add和taobao.promotionmisc.activity.range.remove接口。 &lt;br/&gt;2、使用该接口时需要同时把未做修改的字段值也传入。 &lt;br/&gt;&lt;br/&gt;3、该接口受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
func TaobaoPromotionmiscItemActivityUpdate(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscItemActivityUpdateAPIRequest, resp *promotion.TaobaoPromotionmiscItemActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
