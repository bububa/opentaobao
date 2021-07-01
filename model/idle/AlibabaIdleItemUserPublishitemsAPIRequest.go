package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleItemUserPublishitemsAPIRequest
发布的商品列表 API请求
alibaba.idle.item.user.publishitems

为服务商的卖家提供发布的闲鱼商品列表 */
type AlibabaIdleItemUserPublishitemsAPIRequest struct {
	model.Params
	// 查询参数
	_paramItemPageQuery *ItemPageQuery
}

// New
