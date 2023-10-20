package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaMsfserviceAuditsUpdate 操作改约审批单
// alibaba.msfservice.audits.update
//
// 操作改约审批单
func AlibabaMsfserviceAuditsUpdate(clt *core.SDKClient, req *tmallsc.AlibabaMsfserviceAuditsUpdateAPIRequest, resp *tmallsc.AlibabaMsfserviceAuditsUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
