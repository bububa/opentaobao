package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyItemBatchOperateAPIRequest
批量添加/删除商品和分类的关联关系 API请求
taobao.omniitem.classify.item.batch.operate

批量添加/删除商品和分类的关联关系 */
type TaobaoOmniitemClassifyItemBatchOperateAPIRequest struct {
	model.Params
	// 分类ID
	_classifyIds []int64
	// 需要添加分类关联关系的商品ID
	_addItemIds []int64
	// 需要删除分类关联关系的商品ID
	_deleteItemIds []int64
	// 操作人信息（暂时不填）
	_operator string
}

// New
