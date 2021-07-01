package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSchemaAddAPIRequest
天猫根据规则发布商品 API请求
tmall.item.schema.add

天猫TopSchema发布商品。 */
type TmallItemSchemaAddAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId
	_productId int64
	// 根据tmall.item.add.schema.get生成的商品发布规则入参数据
	_xmlData string
}

// New
