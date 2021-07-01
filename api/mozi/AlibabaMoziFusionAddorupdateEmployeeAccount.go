package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

/* AlibabaMoziFusionAddorupdateEmployeeAccount
添加人员和账号复合接口
alibaba.mozi.fusion.addorupdate.employee.account

添加人员和账号复合接口 */
func AlibabaMoziFusionAddorupdateEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest, session string) (*mozi.AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse, error) {
	var resp mozi.AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
