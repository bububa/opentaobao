package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivitySend 活动权益发放接口
// taobao.promotion.benefit.activity.send
//
// 活动权益发放接口，用于卖家针对活动进行权益发放
func TaobaoPromotionBenefitActivitySend(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivitySendAPIRequest, session string) (*promotion.TaobaoPromotionBenefitActivitySendAPIResponse, error) {
	var resp promotion.TaobaoPromotionBenefitActivitySendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
