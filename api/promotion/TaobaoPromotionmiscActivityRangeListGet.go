package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscActivityRangeListGet 查询活动参与的商品
// taobao.promotionmisc.activity.range.list.get
//
// 查询活动参与的商品
func TaobaoPromotionmiscActivityRangeListGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscActivityRangeListGetAPIRequest, session string) (*promotion.TaobaoPromotionmiscActivityRangeListGetAPIResponse, error) {
	var resp promotion.TaobaoPromotionmiscActivityRangeListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
