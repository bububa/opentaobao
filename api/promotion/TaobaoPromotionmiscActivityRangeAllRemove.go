package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscactivityrangeallremove 清空活动参与的商品
// taobao.promotionmisc.activity.range.all.remove
//
// 清空活动参与的商品
func Taobaopromotionmiscactivityrangeallremove(clt *core.SDKClient, req *promotion.TaobaopromotionmiscactivityrangeallremoveAPIRequest, session string) (*promotion.TaobaopromotionmiscactivityrangeallremoveAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscactivityrangeallremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
