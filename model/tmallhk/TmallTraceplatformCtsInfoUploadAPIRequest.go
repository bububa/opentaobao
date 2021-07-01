package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraceplatformCtsInfoUploadAPIRequest
CTS提交溯源信息 API请求
tmall.traceplatform.cts.info.upload

cts上传溯源信息 */
type TmallTraceplatformCtsInfoUploadAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *CtsInfo
}

// New
