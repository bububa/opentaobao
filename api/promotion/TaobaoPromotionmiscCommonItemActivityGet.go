package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemActivityGet 查询通用单品优惠活动
// taobao.promotionmisc.common.item.activity.get
//
// 查询通用单品优惠活动。
func TaobaoPromotionmiscCommonItemActivityGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityGetAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemActivityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
