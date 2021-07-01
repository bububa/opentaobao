package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupAddAPIRequest
信息流增加单元 API请求
taobao.feedflow.item.adgroup.add

信息流增加单元 */
type TaobaoFeedflowItemAdgroupAddAPIRequest struct {
	model.Params
	// 单元信息
	_adgroup *AdgroupDto
}

// New
