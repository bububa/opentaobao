package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotagTagFindAPIRequest
查询标签接口 API请求
tmall.promotag.tag.find

查询用户创建的所有标签 */
type TmallPromotagTagFindAPIRequest struct {
	model.Params
	// 当前页码
	_pageNo int64
	// 每页显示个数
	_pageSize int64
	// 标签名称，查询时可选项
	_tagName string
	// 标签ID
	_tagId int64
}

// New
