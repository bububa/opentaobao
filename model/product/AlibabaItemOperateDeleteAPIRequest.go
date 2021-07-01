package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemOperateDeleteAPIRequest
商品删除 API请求
alibaba.item.operate.delete

商品删除 */
type AlibabaItemOperateDeleteAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// New
