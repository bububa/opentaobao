package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSchemaMatchAPIRequest
product匹配接口 API请求
tmall.product.schema.match

根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）； */
type TmallProductSchemaMatchAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
	_propvalues string
}

// New
