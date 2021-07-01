package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdzoneListAPIRequest
批量查询可用广告位列表 API请求
taobao.feedflow.item.adzone.list

批量查询可用广告位列表 */
type TaobaoFeedflowItemAdzoneListAPIRequest struct {
	model.Params
	// 广告位查询条件
	_adzoneQuery *AdzoneQueryDto
}

// New
