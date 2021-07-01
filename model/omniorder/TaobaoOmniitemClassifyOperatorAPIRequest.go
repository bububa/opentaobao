package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyOperatorAPIRequest
添加/修改分类 API请求
taobao.omniitem.classify.operator

添加/修改分类 */
type TaobaoOmniitemClassifyOperatorAPIRequest struct {
	model.Params
	// 分类信息
	_category *OmniItemCategoryDetailDto
	// 需要添加的关联关系的商品
	_addItemIds []int64
	// 需要修改的关联关系的商品
	_removeItemIds []int64
	// 操作人信息（暂时不填）
	_operator string
}

// New
