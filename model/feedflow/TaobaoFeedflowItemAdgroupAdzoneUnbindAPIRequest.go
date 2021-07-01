package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest
信息流单元内解绑资源位 API请求
taobao.feedflow.item.adgroup.adzone.unbind

信息流单元内解绑资源位 */
type TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest struct {
	model.Params
	// 广告位id
	_adzoneIdList []int64
	// 单元id
	_adgroupId int64
}

// New
