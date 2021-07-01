package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemUpdateSimpleschemaGetAPIRequest
官网同购编辑商品的get接口 API请求
tmall.item.update.simpleschema.get

官网同购编辑商品的get接口 */
type TmallItemUpdateSimpleschemaGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
