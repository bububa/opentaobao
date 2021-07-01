package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMobilePromotionBenefitActivitySendAPIRequest
手淘专用单用户发放接口 API请求
taobao.mobile.promotion.benefit.activity.send

卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。 */
type TaobaoMobilePromotionBenefitActivitySendAPIRequest struct {
	model.Params
	// 单用户权益发放请求
	_singleBenefitRequest *SingleBenefitRequest
}

// New
