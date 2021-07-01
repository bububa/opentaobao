package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscCommonItemDetailListGetAPIRequest
查询通用单品优惠详情列表 API请求
taobao.promotionmisc.common.item.detail.list.get

查询通用单品优惠详情列表。 */
type TaobaoPromotionmiscCommonItemDetailListGetAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 分页页码，页码从1开始
	_pageNo int64
	// 分页大小，不能超过50
	_pageSize int64
}

// New
