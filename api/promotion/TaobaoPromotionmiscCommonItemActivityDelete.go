package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemActivityDelete 删除通用单品优惠活动
// taobao.promotionmisc.common.item.activity.delete
//
// 删除通用单品优惠活动。
func TaobaoPromotionmiscCommonItemActivityDelete(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
