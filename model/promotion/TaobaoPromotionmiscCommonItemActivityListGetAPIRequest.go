package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscCommonItemActivityListGetAPIRequest
查询通用单品优惠活动列表 API请求
taobao.promotionmisc.common.item.activity.list.get

查询通用单品优惠活动列表。 */
type TaobaoPromotionmiscCommonItemActivityListGetAPIRequest struct {
	model.Params
	// 分页页码，页码从1开始
	_pageNo int64
	// 分页大小，不能超过50
	_pageSize int64
}

// New
