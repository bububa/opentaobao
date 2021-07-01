package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSchemaAddAPIRequest
使用Schema文件发布一个产品 API请求
tmall.product.schema.add

Schema体系发布一个产品 */
type TmallProductSchemaAddAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 品牌ID
	_brandId int64
	// 根据tmall.product.add.schema.get生成的产品发布规则入参数据
	_xmlData string
}

// New
