package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscitemactivityadd 创建无条件单品优惠活动
// taobao.promotionmisc.item.activity.add
//
// 创建无条件单品优惠活动。1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。&lt;br/&gt;2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。&lt;br/&gt;3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
func Taobaopromotionmiscitemactivityadd(clt *core.SDKClient, req *promotion.TaobaopromotionmiscitemactivityaddAPIRequest, session string) (*promotion.TaobaopromotionmiscitemactivityaddAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscitemactivityaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
