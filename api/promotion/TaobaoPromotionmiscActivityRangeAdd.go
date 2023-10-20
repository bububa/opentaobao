package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscactivityrangeadd 增加活动参与的商品
// taobao.promotionmisc.activity.range.add
//
// 增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
func Taobaopromotionmiscactivityrangeadd(clt *core.SDKClient, req *promotion.TaobaopromotionmiscactivityrangeaddAPIRequest, session string) (*promotion.TaobaopromotionmiscactivityrangeaddAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscactivityrangeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
