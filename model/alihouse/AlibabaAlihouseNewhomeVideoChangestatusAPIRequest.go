package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeVideoChangestatusAPIRequest
视频草稿状态更新 API请求
alibaba.alihouse.newhome.video.changestatus

视频草稿状态更新 */
type AlibabaAlihouseNewhomeVideoChangestatusAPIRequest struct {
	model.Params
	// 外部视频id
	_outerId string
	// 0 失效 1 有效
	_status int64
}

// New
