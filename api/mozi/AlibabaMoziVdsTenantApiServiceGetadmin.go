package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozivdstenantapiservicegetadmin 获取员工租户管理员信息（查询员工是否为租户管理员）
// alibaba.mozi.vds.tenant.api.service.getadmin
//
// 获取员工租户管理员信息（查询员工是否为租户管理员）
func Alibabamozivdstenantapiservicegetadmin(clt *core.SDKClient, req *mozi.AlibabamozivdstenantapiservicegetadminAPIRequest, session string) (*mozi.AlibabamozivdstenantapiservicegetadminAPIResponse, error) {
	var resp mozi.AlibabamozivdstenantapiservicegetadminAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
