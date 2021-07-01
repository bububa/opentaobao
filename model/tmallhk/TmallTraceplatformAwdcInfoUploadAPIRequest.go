package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraceplatformAwdcInfoUploadAPIRequest
AWDC提交溯源信息 API请求
tmall.traceplatform.awdc.info.upload

天猫溯源-AWDC-上传溯源信息 */
type TmallTraceplatformAwdcInfoUploadAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *AwdcInfo
}

// New
