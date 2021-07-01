package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSimpleschemaUpdateAPIRequest
天猫简化编辑商品 API请求
tmall.item.simpleschema.update

国外大商家天猫简化编辑商品 */
type TmallItemSimpleschemaUpdateAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 编辑商品时提交的xml信息
	_schemaXmlFields string
}

// New
