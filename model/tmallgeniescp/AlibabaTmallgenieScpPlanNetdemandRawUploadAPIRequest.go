package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest
二级物料净需求回传(TL+1) API请求
alibaba.tmallgenie.scp.plan.netdemand.raw.upload

二级物料净需求回传(TL+1) */
type AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRawRequest *NetDemandRawRequest
}

// New
