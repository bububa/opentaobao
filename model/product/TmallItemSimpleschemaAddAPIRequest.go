package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSimpleschemaAddAPIRequest
天猫简化发布商品 API请求
tmall.item.simpleschema.add

天猫简化版schema发布商品。 */
type TmallItemSimpleschemaAddAPIRequest struct {
	model.Params
	// 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
	_schemaXmlFields string
}

// New
