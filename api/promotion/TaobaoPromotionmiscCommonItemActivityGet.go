package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

/* TaobaoPromotionmiscCommonItemActivityGet
查询通用单品优惠活动
taobao.promotionmisc.common.item.activity.get

查询通用单品优惠活动。 */
func TaobaoPromotionmiscCommonItemActivityGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityGetAPIRequest, session string) (*promotion.TaobaoPromotionmiscCommonItemActivityGetAPIResponse, error) {
	var resp promotion.TaobaoPromotionmiscCommonItemActivityGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
