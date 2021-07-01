package aliexpress

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressFeedPostPublishAPIRequest
同步帖子 API请求
aliexpress.feed.post.publish

站外平台同步帖子至AE FEED域 */
type AliexpressFeedPostPublishAPIRequest struct {
	model.Params
	// 站外导入内容请求参数
	_offsitePublishPostEntity *OffsitePublishPostEntity
}

// New
