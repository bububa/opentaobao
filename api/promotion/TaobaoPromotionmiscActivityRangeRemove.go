package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscactivityrangeremove 去除活动参与的商品
// taobao.promotionmisc.activity.range.remove
//
// 去除活动参与的商品
func Taobaopromotionmiscactivityrangeremove(clt *core.SDKClient, req *promotion.TaobaopromotionmiscactivityrangeremoveAPIRequest, session string) (*promotion.TaobaopromotionmiscactivityrangeremoveAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscactivityrangeremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
