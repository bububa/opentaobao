package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductMatchSchemaGetAPIRequest
获取匹配产品规则 API请求
tmall.product.match.schema.get

ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则 */
type TmallProductMatchSchemaGetAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
}

// New
