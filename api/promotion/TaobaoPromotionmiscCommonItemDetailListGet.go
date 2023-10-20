package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemDetailListGet 查询通用单品优惠详情列表
// taobao.promotionmisc.common.item.detail.list.get
//
// 查询通用单品优惠详情列表。
func TaobaoPromotionmiscCommonItemDetailListGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemDetailListGetAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemDetailListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
