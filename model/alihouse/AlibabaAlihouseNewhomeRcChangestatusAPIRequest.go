package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeRcChangestatusAPIRequest
图文草稿状态更新 API请求
alibaba.alihouse.newhome.rc.changestatus

图文草稿状态更新 */
type AlibabaAlihouseNewhomeRcChangestatusAPIRequest struct {
	model.Params
	// 外部图文id
	_outerId string
	// 0 失效 1 有效
	_status int64
}

// New
