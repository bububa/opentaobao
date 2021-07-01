package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscMjsActivityGetAPIRequest
查询满就送活动 API请求
taobao.promotionmisc.mjs.activity.get

查询满就送活动 */
type TaobaoPromotionmiscMjsActivityGetAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// New
