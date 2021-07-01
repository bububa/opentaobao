package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentItemEditAPIRequest
租赁商品编辑 API请求
alibaba.idle.rent.item.edit

发布闲鱼租赁商品 */
type AlibabaIdleRentItemEditAPIRequest struct {
	model.Params
	// 商品信息
	_paramRentItemDTO *RentItemDto
}

// New
