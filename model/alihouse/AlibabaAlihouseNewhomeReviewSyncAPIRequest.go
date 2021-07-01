package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeReviewSyncAPIRequest
天猫好房楼盘评测同步 API请求
alibaba.alihouse.newhome.review.sync

接受楼盘测评图文信息 */
type AlibabaAlihouseNewhomeReviewSyncAPIRequest struct {
	model.Params
	// 测评草稿信息
	_review *ProjectReviewDraftDto
}

// New
