package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpActivitiesGetAPIRequest
查询活动列表 API请求
taobao.ump.activities.get

查询活动列表 */
type TaobaoUmpActivitiesGetAPIRequest struct {
	model.Params
	// 工具id
	_toolId int64
	// 分页的页码
	_pageNo int64
	// 每页的最大条数
	_pageSize int64
}

// New
