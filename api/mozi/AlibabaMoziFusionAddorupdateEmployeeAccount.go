package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionAddorupdateEmployeeAccount 添加人员和账号复合接口
// alibaba.mozi.fusion.addorupdate.employee.account
//
// 添加人员和账号复合接口
func AlibabaMoziFusionAddorupdateEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest, resp *mozi.AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
