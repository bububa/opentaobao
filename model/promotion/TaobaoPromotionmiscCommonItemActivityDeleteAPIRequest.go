package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest
删除通用单品优惠活动 API请求
taobao.promotionmisc.common.item.activity.delete

删除通用单品优惠活动。 */
type TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
}

// New
