package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscCommonItemActivityGetAPIRequest
查询通用单品优惠活动 API请求
taobao.promotionmisc.common.item.activity.get

查询通用单品优惠活动。 */
type TaobaoPromotionmiscCommonItemActivityGetAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
}

// New
