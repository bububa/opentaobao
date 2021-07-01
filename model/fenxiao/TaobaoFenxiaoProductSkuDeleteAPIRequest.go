package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductSkuDeleteAPIRequest
产品SKU删除接口 API请求
taobao.fenxiao.product.sku.delete

根据sku properties删除sku数据 */
type TaobaoFenxiaoProductSkuDeleteAPIRequest struct {
	model.Params
	// 产品id
	_productId int64
	// sku属性
	_properties string
}

// New
