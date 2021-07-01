package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemOperateDownshelfAPIRequest
商品下架 API请求
alibaba.item.operate.downshelf

商品下架 */
type AlibabaItemOperateDownshelfAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// New
