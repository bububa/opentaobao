package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpRangeGetAPIRequest
查询活动范围 API请求
taobao.ump.range.get

查询某个卖家所有参加或者不参加某项活动的物品 */
type TaobaoUmpRangeGetAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// New
