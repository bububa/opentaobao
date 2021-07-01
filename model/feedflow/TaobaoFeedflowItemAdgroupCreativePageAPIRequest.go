package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupCreativePageAPIRequest
信息流单元下查看创意 API请求
taobao.feedflow.item.adgroup.creative.page

信息流单元下查看创意 */
type TaobaoFeedflowItemAdgroupCreativePageAPIRequest struct {
	model.Params
	// 绑定查询条件
	_creativeBindQuery *CreativeBindQueryDto
}

// New
