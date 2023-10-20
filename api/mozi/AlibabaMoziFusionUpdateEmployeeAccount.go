package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionUpdateEmployeeAccount 更新人员和账号属性
// alibaba.mozi.fusion.update.employee.account
//
// 更新人员和账号基本属性
func AlibabaMoziFusionUpdateEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionUpdateEmployeeAccountAPIRequest, resp *mozi.AlibabaMoziFusionUpdateEmployeeAccountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
