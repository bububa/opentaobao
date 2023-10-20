package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionReentryEmployeeAccount 重新入职并且重新启用账号
// alibaba.mozi.fusion.reentry.employee.account
//
// 重新入职并且重新启用账号
func AlibabaMoziFusionReentryEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionReentryEmployeeAccountAPIRequest, resp *mozi.AlibabaMoziFusionReentryEmployeeAccountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
