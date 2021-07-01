package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpActivityGetAPIRequest
查询营销活动 API请求
taobao.ump.activity.get

查询营销活动 */
type TaobaoUmpActivityGetAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// New
