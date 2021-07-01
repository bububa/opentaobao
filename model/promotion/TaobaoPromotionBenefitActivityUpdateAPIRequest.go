package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionBenefitActivityUpdateAPIRequest
修改关联的活动权益 API请求
taobao.promotion.benefit.activity.update

修改卖家活动中关联的对应的权益。 */
type TaobaoPromotionBenefitActivityUpdateAPIRequest struct {
	model.Params
	// 修改关联的权益的活动请求
	_updateRequest *UpdateBenefitActivityRequest
}

// New
