package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemDapeiTemplateQueryAPIRequest
搭配查询接口 API请求
tmall.item.dapei.template.query

根据条件获取搭配内容 */
type TmallItemDapeiTemplateQueryAPIRequest struct {
	model.Params
	// 搭配标题
	_title string
	// 页码
	_pageIndex int64
	// 分页大小
	_pageSize int64
}

// New
