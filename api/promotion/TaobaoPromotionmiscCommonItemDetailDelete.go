package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemDetailDelete 删除通用单品优惠详情
// taobao.promotionmisc.common.item.detail.delete
//
// 删除通用单品优惠详情。
func TaobaoPromotionmiscCommonItemDetailDelete(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
