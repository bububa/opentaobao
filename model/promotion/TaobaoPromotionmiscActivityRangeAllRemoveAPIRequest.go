package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest
清空活动参与的商品 API请求
taobao.promotionmisc.activity.range.all.remove

清空活动参与的商品 */
type TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// New
