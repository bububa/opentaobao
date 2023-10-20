package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoMobilePromotionBenefitActivitySend 手淘专用单用户发放接口
// taobao.mobile.promotion.benefit.activity.send
//
// 卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
func TaobaoMobilePromotionBenefitActivitySend(clt *core.SDKClient, req *promotion.TaobaoMobilePromotionBenefitActivitySendAPIRequest, resp *promotion.TaobaoMobilePromotionBenefitActivitySendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
