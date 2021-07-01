package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscItemActivityListGetAPIRequest
查询无条件单品优惠活动列表 API请求
taobao.promotionmisc.item.activity.list.get

查询无条件单品优惠活动列表 */
type TaobaoPromotionmiscItemActivityListGetAPIRequest struct {
	model.Params
	// 页码。
	_pageNo int64
	// 每页记录数，最大支持50 。
	_pageSize int64
}

// New
