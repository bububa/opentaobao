package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSchemaUpdateAPIRequest
产品更新接口 API请求
tmall.product.schema.update

产品更新接口 */
type TmallProductSchemaUpdateAPIRequest struct {
	model.Params
	// 根据tmall.product.update.schema.get生成的产品更新规则入参数据
	_xmlData string
	// 产品编号
	_productId int64
}

// New
