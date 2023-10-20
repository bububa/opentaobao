package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscactivityrangelistget 查询活动参与的商品
// taobao.promotionmisc.activity.range.list.get
//
// 查询活动参与的商品
func Taobaopromotionmiscactivityrangelistget(clt *core.SDKClient, req *promotion.TaobaopromotionmiscactivityrangelistgetAPIRequest, session string) (*promotion.TaobaopromotionmiscactivityrangelistgetAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscactivityrangelistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
