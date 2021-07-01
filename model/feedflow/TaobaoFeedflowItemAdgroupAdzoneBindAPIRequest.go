package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest
信息流单元内绑定资源位 API请求
taobao.feedflow.item.adgroup.adzone.bind

信息流单元内绑定资源位 */
type TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest struct {
	model.Params
	// 新增的绑定资源位
	_bindAdzoneList []AdzoneBindDto
	// 单元id
	_adgroupId int64
}

// New
