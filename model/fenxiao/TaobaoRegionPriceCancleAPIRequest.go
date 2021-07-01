package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRegionPriceCancleAPIRequest
取消区域价格 API请求
taobao.region.price.cancle

取消区域价格 */
type TaobaoRegionPriceCancleAPIRequest struct {
	model.Params
	// 商品
	_itemId int64
	// 无sku传0
	_skuId int64
}

// New
