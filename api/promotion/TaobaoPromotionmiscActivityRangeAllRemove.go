package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscActivityRangeAllRemove 清空活动参与的商品
// taobao.promotionmisc.activity.range.all.remove
//
// 清空活动参与的商品
func TaobaoPromotionmiscActivityRangeAllRemove(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest, resp *promotion.TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
