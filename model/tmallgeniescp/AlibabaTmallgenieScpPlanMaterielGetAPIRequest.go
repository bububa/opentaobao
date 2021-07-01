package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanMaterielGetAPIRequest
1-IBP同步物料接口 API请求
alibaba.tmallgenie.scp.plan.materiel.get

IBP同步物料接口 */
type AlibabaTmallgenieScpPlanMaterielGetAPIRequest struct {
	model.Params
	// 扩展字段
	_requestExtendJson string
}

// New
