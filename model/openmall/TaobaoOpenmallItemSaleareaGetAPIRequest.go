package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallItemSaleareaGetAPIRequest
查询商品可售区域 API请求
taobao.openmall.item.salearea.get

获取openmall商品的可售区域 */
type TaobaoOpenmallItemSaleareaGetAPIRequest struct {
	model.Params
	// 商品SKU
	_skuIds string
	// 商品ID
	_itemId int64
}

// New
