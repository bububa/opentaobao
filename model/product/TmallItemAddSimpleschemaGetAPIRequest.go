package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemAddSimpleschemaGetAPIRequest
天猫发布商品规则获取 API请求
tmall.item.add.simpleschema.get

通过商家信息获取商品发布字段和规则。 */
type TmallItemAddSimpleschemaGetAPIRequest struct {
	model.Params
}

// New
