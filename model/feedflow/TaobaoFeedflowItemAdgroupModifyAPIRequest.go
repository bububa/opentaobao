package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupModifyAPIRequest
信息流单元修改 API请求
taobao.feedflow.item.adgroup.modify

信息流单元修改 */
type TaobaoFeedflowItemAdgroupModifyAPIRequest struct {
	model.Params
	// 单元信息
	_adgroup *AdgroupDto
}

// New
