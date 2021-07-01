package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* AlibabaTmallgenieScpPlanHistoryPoGet
【已废除】11-同步历史所有的po单
alibaba.tmallgenie.scp.plan.history.po.get

同步历史po单 */
func AlibabaTmallgenieScpPlanHistoryPoGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanHistoryPoGetAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanHistoryPoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
