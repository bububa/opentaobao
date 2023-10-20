package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscMjsActivityAdd 创建满就送活动
// taobao.promotionmisc.mjs.activity.add
//
// 创建满就送活动。&lt;br/&gt;1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。 2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。 3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
func TaobaoPromotionmiscMjsActivityAdd(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscMjsActivityAddAPIRequest, resp *promotion.TaobaoPromotionmiscMjsActivityAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
