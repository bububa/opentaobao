package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpDetailsGetAPIRequest
查询活动详情列表 API请求
taobao.ump.details.get

分页查询优惠详情列表 */
type TaobaoUmpDetailsGetAPIRequest struct {
	model.Params
	// 营销活动id
	_actId int64
	// 分页的页码
	_pageNo int64
	// 每页的最大条数
	_pageSize int64
}

// New
