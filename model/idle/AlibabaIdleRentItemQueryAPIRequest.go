package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentItemQueryAPIRequest
查询租赁商品信息 API请求
alibaba.idle.rent.item.query

查询租赁商品信息 */
type AlibabaIdleRentItemQueryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
