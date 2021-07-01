package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentItemAddAPIRequest
租赁商品发布 API请求
alibaba.idle.rent.item.add

发布闲鱼租赁商品 */
type AlibabaIdleRentItemAddAPIRequest struct {
	model.Params
	// 商品信息
	_paramRentItemDTO *RentItemDto
}

// New
