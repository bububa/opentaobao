package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionDimissionEmployeeAccount 人员离职
// alibaba.mozi.fusion.dimission.employee.account
//
// 人员离职并且回收账号
func AlibabaMoziFusionDimissionEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionDimissionEmployeeAccountAPIRequest, resp *mozi.AlibabaMoziFusionDimissionEmployeeAccountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
