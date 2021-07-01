package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

/* AlibabaMoziFusionCreateEmployeeAccount
创建MOZI自建人员和账号
alibaba.mozi.fusion.create.employee.account

创建MOZI自建人员和账号 */
func AlibabaMoziFusionCreateEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionCreateEmployeeAccountAPIRequest, session string) (*mozi.AlibabaMoziFusionCreateEmployeeAccountAPIResponse, error) {
	var resp mozi.AlibabaMoziFusionCreateEmployeeAccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
