package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziVdsTenantApiServiceTenantbyid 按租户ID查询租户信息
// alibaba.mozi.vds.tenant.api.service.tenantbyid
//
// 按租户ID查询租户信息
func AlibabaMoziVdsTenantApiServiceTenantbyid(clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest, session string) (*mozi.AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse, error) {
	var resp mozi.AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
