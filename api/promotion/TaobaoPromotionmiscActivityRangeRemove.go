package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscActivityRangeRemove 去除活动参与的商品
// taobao.promotionmisc.activity.range.remove
//
// 去除活动参与的商品
func TaobaoPromotionmiscActivityRangeRemove(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscActivityRangeRemoveAPIRequest, session string) (*promotion.TaobaoPromotionmiscActivityRangeRemoveAPIResponse, error) {
	var resp promotion.TaobaoPromotionmiscActivityRangeRemoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
