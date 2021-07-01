package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeLineSyncAPIRequest
环线数据同步 API请求
alibaba.alihouse.newhome.line.sync

环线数据同步 */
type AlibabaAlihouseNewhomeLineSyncAPIRequest struct {
	model.Params
	// 环线入参
	_baseLoopLineDto *BaseLoopLineDto
}

// New
