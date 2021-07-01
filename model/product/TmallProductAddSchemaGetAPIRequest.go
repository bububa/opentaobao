package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductAddSchemaGetAPIRequest
产品发布规则获取接口 API请求
tmall.product.add.schema.get

获取用户发布产品的规则 */
type TmallProductAddSchemaGetAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 品牌ID
	_brandId int64
}

// New
