package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCreativeDeleteAPIRequest
信息流删除创意 API请求
taobao.feedflow.item.creative.delete

信息流删除创意 */
type TaobaoFeedflowItemCreativeDeleteAPIRequest struct {
	model.Params
	// 创意id列表
	_creativeIdList []int64
}

// New
