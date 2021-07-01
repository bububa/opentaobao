package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemItemPageAPIRequest
信息流查看商品列表 API请求
taobao.feedflow.item.item.page

信息流查看商品列表 */
type TaobaoFeedflowItemItemPageAPIRequest struct {
	model.Params
	// 查询条件
	_itemQuery *ItemQueryDto
}

// New
