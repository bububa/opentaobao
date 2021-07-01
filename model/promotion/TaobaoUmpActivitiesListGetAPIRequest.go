package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpActivitiesListGetAPIRequest
营销活动列表查询 API请求
taobao.ump.activities.list.get

按照营销活动id的列表ids，查询对应的营销活动列表。 */
type TaobaoUmpActivitiesListGetAPIRequest struct {
	model.Params
	// 营销活动id列表。
	_ids []int64
}

// New
