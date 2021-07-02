package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionDimissionEmployeeAccount 人员离职
// alibaba.mozi.fusion.dimission.employee.account
//
// 人员离职并且回收账号
func AlibabaMoziFusionDimissionEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionDimissionEmployeeAccountAPIRequest, session string) (*mozi.AlibabaMoziFusionDimissionEmployeeAccountAPIResponse, error) {
	var resp mozi.AlibabaMoziFusionDimissionEmployeeAccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
