package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscActivityRangeAddAPIRequest
增加活动参与的商品 API请求
taobao.promotionmisc.activity.range.add

增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。 */
type TaobaoPromotionmiscActivityRangeAddAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
	// 商品id,多个id用逗号隔开，一次不超过50个。
	_ids string
}

// New
