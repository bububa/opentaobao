package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozivdstenantapiservicepagesubadmins 分页查询租户子管理员
// alibaba.mozi.vds.tenant.api.service.pagesubadmins
//
// 分页查询租户子管理员
func Alibabamozivdstenantapiservicepagesubadmins(clt *core.SDKClient, req *mozi.AlibabamozivdstenantapiservicepagesubadminsAPIRequest, session string) (*mozi.AlibabamozivdstenantapiservicepagesubadminsAPIResponse, error) {
	var resp mozi.AlibabamozivdstenantapiservicepagesubadminsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
