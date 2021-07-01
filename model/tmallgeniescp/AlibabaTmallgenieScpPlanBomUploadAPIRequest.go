package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanBomUploadAPIRequest
计划BOM同步 API请求
alibaba.tmallgenie.scp.plan.bom.upload

计划BOM同步 */
type AlibabaTmallgenieScpPlanBomUploadAPIRequest struct {
	model.Params
	// 对象
	_pbomRequest *AbstractRequest
}

// New
