package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyStoreBatchOperateAPIRequest
批量添加/删除门店和分类的关联关系 API请求
taobao.omniitem.classify.store.batch.operate

批量添加/删除门店和分类的关联关系 */
type TaobaoOmniitemClassifyStoreBatchOperateAPIRequest struct {
	model.Params
	// 商家门店ID
	_storeIds []int64
	// 需要添加的分类ID
	_addCategoryIds []int64
	// 需要删除的分类ID
	_removeCategoryIds []int64
	// 操作信息（暂时不填）
	_operator string
}

// New
