package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscMjsActivityListGetAPIRequest
查询满就送活动列表 API请求
taobao.promotionmisc.mjs.activity.list.get

查询满就送活动列表。注意，该接口的返回值中，只包含活动的主要信息，如activity_id，name，description，start_time，end_time，type，participate_range。优惠的详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。 */
type TaobaoPromotionmiscMjsActivityListGetAPIRequest struct {
	model.Params
	// 活动类型： 1表示商品级别的活动；2表示店铺级别的活动。
	_activityType int64
	// 页码。
	_pageNo int64
	// 每页记录数，最大支持50 。
	_pageSize int64
}

// New
