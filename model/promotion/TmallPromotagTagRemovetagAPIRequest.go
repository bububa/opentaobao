package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotagTagRemovetagAPIRequest
删除标签定义 API请求
tmall.promotag.tag.removetag

用于删除标签定义，但是要确保目前该标签没有人群在使用。 */
type TmallPromotagTagRemovetagAPIRequest struct {
	model.Params
	// 需要删除的标签id
	_tagId int64
}

// New
