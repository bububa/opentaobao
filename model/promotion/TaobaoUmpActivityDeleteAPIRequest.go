package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpActivityDeleteAPIRequest
删除营销活动 API请求
taobao.ump.activity.delete

删除营销活动。对应的活动详情等将会被全部删除。 */
type TaobaoUmpActivityDeleteAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// New
