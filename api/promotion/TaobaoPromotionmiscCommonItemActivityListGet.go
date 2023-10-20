package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemActivityListGet 查询通用单品优惠活动列表
// taobao.promotionmisc.common.item.activity.list.get
//
// 查询通用单品优惠活动列表。
func TaobaoPromotionmiscCommonItemActivityListGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityListGetAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemActivityListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
