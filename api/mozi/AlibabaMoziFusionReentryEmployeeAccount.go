package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionReentryEmployeeAccount 重新入职并且重新启用账号
// alibaba.mozi.fusion.reentry.employee.account
//
// 重新入职并且重新启用账号
func AlibabaMoziFusionReentryEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionReentryEmployeeAccountAPIRequest, session string) (*mozi.AlibabaMoziFusionReentryEmployeeAccountAPIResponse, error) {
	var resp mozi.AlibabaMoziFusionReentryEmployeeAccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
