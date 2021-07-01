package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemCombineGetAPIRequest
组合商品获取接口 API请求
tmall.item.combine.get

查询组合商品的SKU信息 */
type TmallItemCombineGetAPIRequest struct {
	model.Params
	// 组合商品ID
	_itemId int64
}

// New
