package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductUpdateSchemaGetAPIRequest
产品更新规则获取接口 API请求
tmall.product.update.schema.get

获取用户更新产品的规则 */
type TmallProductUpdateSchemaGetAPIRequest struct {
	model.Params
	// 产品编号
	_productId int64
}

// New
