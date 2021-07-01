package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductSkusGetAPIRequest
SKU查询接口 API请求
taobao.fenxiao.product.skus.get

产品sku查询 */
type TaobaoFenxiaoProductSkusGetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// New
