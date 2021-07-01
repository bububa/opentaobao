package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSchemaGetAPIRequest
产品信息获取schema获取 API请求
tmall.product.schema.get

产品信息获取接口schema形式返回 */
type TmallProductSchemaGetAPIRequest struct {
	model.Params
	// 产品编号
	_productId int64
}

// New
