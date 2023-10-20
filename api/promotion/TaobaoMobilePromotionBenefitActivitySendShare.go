package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoMobilePromotionBenefitActivitySendShare 手淘专用单用户发放接口
// taobao.mobile.promotion.benefit.activity.send.share
//
// 卖家活动中需要通过该API来发放对应的权益。手淘专用、验证分享链路。
func TaobaoMobilePromotionBenefitActivitySendShare(clt *core.SDKClient, req *promotion.TaobaoMobilePromotionBenefitActivitySendShareAPIRequest, resp *promotion.TaobaoMobilePromotionBenefitActivitySendShareAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
