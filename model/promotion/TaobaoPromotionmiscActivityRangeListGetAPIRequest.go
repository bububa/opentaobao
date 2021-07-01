package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscActivityRangeListGetAPIRequest
查询活动参与的商品 API请求
taobao.promotionmisc.activity.range.list.get

查询活动参与的商品 */
type TaobaoPromotionmiscActivityRangeListGetAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
}

// New
