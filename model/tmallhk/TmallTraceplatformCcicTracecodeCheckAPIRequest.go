package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraceplatformCcicTracecodeCheckAPIRequest
ccic校验溯源码 API请求
tmall.traceplatform.ccic.tracecode.check

天猫国际溯源业务，需要将溯源码校验的功能输出到ccic官方主页中以增强溯源码的可信度，故需要提供api给ccic使用以校验溯源码的正确性。 */
type TmallTraceplatformCcicTracecodeCheckAPIRequest struct {
	model.Params
	// 15为溯源短码，必选
	_shortTracecode string
	// 6位暗码，必选
	_hideCode string
}

// New
