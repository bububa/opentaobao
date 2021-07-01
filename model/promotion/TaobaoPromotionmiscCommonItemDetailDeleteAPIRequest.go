package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest
删除通用单品优惠详情 API请求
taobao.promotionmisc.common.item.detail.delete

删除通用单品优惠详情。 */
type TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 优惠详情ID
	_detailId int64
}

// New
