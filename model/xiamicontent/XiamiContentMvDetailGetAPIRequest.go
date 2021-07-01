package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentMvDetailGetAPIRequest
获取mv详情 API请求
xiami.content.mv.detail.get

获取mv详情 */
type XiamiContentMvDetailGetAPIRequest struct {
	model.Params
	// mvId
	_mvIds []int64
}

// New
