package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

/* AlibabaMoziFusionUpdateEmployeeAccount
更新人员和账号属性
alibaba.mozi.fusion.update.employee.account

更新人员和账号基本属性 */
func AlibabaMoziFusionUpdateEmployeeAccount(clt *core.SDKClient, req *mozi.AlibabaMoziFusionUpdateEmployeeAccountAPIRequest, session string) (*mozi.AlibabaMoziFusionUpdateEmployeeAccountAPIResponse, error) {
	var resp mozi.AlibabaMoziFusionUpdateEmployeeAccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
