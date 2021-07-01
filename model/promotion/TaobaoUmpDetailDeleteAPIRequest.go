package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpDetailDeleteAPIRequest
删除活动详情 API请求
taobao.ump.detail.delete

删除活动详情 */
type TaobaoUmpDetailDeleteAPIRequest struct {
	model.Params
	// 活动详情id
	_detailId int64
}

// New
