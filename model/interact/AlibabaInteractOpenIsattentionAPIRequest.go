package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractOpenIsattentionAPIRequest
判断用户是否收藏某个店铺 API请求
alibaba.interact.open.isattention

判断用户是否收藏某个店铺 */
type AlibabaInteractOpenIsattentionAPIRequest struct {
	model.Params
	// 1
	_param0 int64
}

// New
