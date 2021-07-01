package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionBenefitActivityTimeUpdateAPIRequest
更新关联活动有效时间 API请求
taobao.promotion.benefit.activity.time.update

更新关联权益的活动有效时间 */
type TaobaoPromotionBenefitActivityTimeUpdateAPIRequest struct {
	model.Params
	// ISV活动关联权益后获得的关联ID
	_relationId int64
	// 活动的开始时间
	_startTime string
	// 活动的i结束时间
	_endTime string
}

// New
