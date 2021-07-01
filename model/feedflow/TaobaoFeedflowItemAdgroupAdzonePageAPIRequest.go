package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupAdzonePageAPIRequest
信息流单元下查看绑定资源位 API请求
taobao.feedflow.item.adgroup.adzone.page

信息流单元下查看绑定资源位 */
type TaobaoFeedflowItemAdgroupAdzonePageAPIRequest struct {
	model.Params
	// 查询条件
	_adzoneBindQuery *AdzoneBindQueryDto
}

// New
