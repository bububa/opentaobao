package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyStoreQueryAPIRequest
根据门店查分类信息 API请求
taobao.omniitem.classify.store.query

根据门店查分类信息 */
type TaobaoOmniitemClassifyStoreQueryAPIRequest struct {
	model.Params
	// 商户的门店ID
	_storeId int64
	// 页码
	_pageNum int64
	// 每页大小
	_pageSize int64
}

// New
