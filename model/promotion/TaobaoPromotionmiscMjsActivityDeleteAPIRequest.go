package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscMjsActivityDeleteAPIRequest
删除满就送活动 API请求
taobao.promotionmisc.mjs.activity.delete

删除满就送活动 */
type TaobaoPromotionmiscMjsActivityDeleteAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// New
