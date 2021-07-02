package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemActivityDelete 删除通用单品优惠活动
// taobao.promotionmisc.common.item.activity.delete
//
// 删除通用单品优惠活动。
func TaobaoPromotionmiscCommonItemActivityDelete(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest, session string) (*promotion.TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse, error) {
	var resp promotion.TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
