package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscActivityRangeRemoveAPIRequest
去除活动参与的商品 API请求
taobao.promotionmisc.activity.range.remove

去除活动参与的商品 */
type TaobaoPromotionmiscActivityRangeRemoveAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
	// 商品id,多个id用逗号隔开。
	_ids string
}

// New
