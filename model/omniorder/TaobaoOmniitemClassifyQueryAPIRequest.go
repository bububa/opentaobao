package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyQueryAPIRequest
查询分类信息 API请求
taobao.omniitem.classify.query

通过查询关键字，分页查询分类信息 */
type TaobaoOmniitemClassifyQueryAPIRequest struct {
	model.Params
	// 查询关键词
	_keyword string
	// 页码
	_pageNum int64
	// 每页大小
	_pageSize int64
}

// New
