package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoMarketingPromotionKfc 定向优惠活动名称与描述违禁词检查
// taobao.marketing.promotion.kfc
//
// 活动名称与描述违禁词检查
func TaobaoMarketingPromotionKfc(clt *core.SDKClient, req *promotion.TaobaoMarketingPromotionKfcAPIRequest, session string) (*promotion.TaobaoMarketingPromotionKfcAPIResponse, error) {
	var resp promotion.TaobaoMarketingPromotionKfcAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
