package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupPageAPIRequest
查询单元列表 API请求
taobao.feedflow.item.adgroup.page

通过计划id查询单元信息 */
type TaobaoFeedflowItemAdgroupPageAPIRequest struct {
	model.Params
	// 系统自动生成
	_adgroupQuery *AdgroupQueryDto
}

// New
