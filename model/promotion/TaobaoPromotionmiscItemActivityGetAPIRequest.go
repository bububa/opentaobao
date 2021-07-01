package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscItemActivityGetAPIRequest
查询无条件单品优惠活动 API请求
taobao.promotionmisc.item.activity.get

查询无条件单品优惠活动 */
type TaobaoPromotionmiscItemActivityGetAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// New
