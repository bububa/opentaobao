package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionBenefitActivityDeleteAPIRequest
删除关联的活动权益 API请求
taobao.promotion.benefit.activity.delete

删除关联的活动权益 */
type TaobaoPromotionBenefitActivityDeleteAPIRequest struct {
	model.Params
	// ISV活动关联权益后获得的关联ID
	_relationId int64
}

// New
