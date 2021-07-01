package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyItemQueryAPIRequest
根据分类查商品信息 API请求
taobao.omniitem.classify.item.query

商家根据分类查商品 */
type TaobaoOmniitemClassifyItemQueryAPIRequest struct {
	model.Params
	// 分类ID
	_classifyId int64
	// 页码
	_pageNum int64
	// 每页大小
	_pageSize int64
}

// New
