package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscMjsActivityGet 查询满就送活动
// taobao.promotionmisc.mjs.activity.get
//
// 查询满就送活动
func TaobaoPromotionmiscMjsActivityGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscMjsActivityGetAPIRequest, session string) (*promotion.TaobaoPromotionmiscMjsActivityGetAPIResponse, error) {
	var resp promotion.TaobaoPromotionmiscMjsActivityGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
