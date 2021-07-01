package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemOperateUpshelfAPIRequest
商品上架 API请求
alibaba.item.operate.upshelf

商品上架 */
type AlibabaItemOperateUpshelfAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品库存
	_quantity int64
}

// New
