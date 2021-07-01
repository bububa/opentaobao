package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscItemActivityDeleteAPIRequest
删除无条件单品优惠活动 API请求
taobao.promotionmisc.item.activity.delete

删除无条件单品优惠活动 */
type TaobaoPromotionmiscItemActivityDeleteAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// New
