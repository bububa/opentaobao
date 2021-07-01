package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeReviewChangestatusAPIRequest
楼盘测评草稿状态同步 API请求
alibaba.alihouse.newhome.review.changestatus

楼盘测评草稿状态更新 */
type AlibabaAlihouseNewhomeReviewChangestatusAPIRequest struct {
	model.Params
	// 外部测评id
	_outerId string
	// 0 失效 1 有效
	_status int64
}

// New
