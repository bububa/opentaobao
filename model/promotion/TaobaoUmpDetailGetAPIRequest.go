package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpDetailGetAPIRequest
查询活动详情 API请求
taobao.ump.detail.get

查询活动详情 */
type TaobaoUmpDetailGetAPIRequest struct {
	model.Params
	// 活动详情的id
	_detailId int64
}

// New
