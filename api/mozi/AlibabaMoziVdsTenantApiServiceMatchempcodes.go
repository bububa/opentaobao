package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziVdsTenantApiServiceMatchempcodes 校验组-员工是否匹配
// alibaba.mozi.vds.tenant.api.service.matchempcodes
//
// 校验组-员工是否匹配
func AlibabaMoziVdsTenantApiServiceMatchempcodes(clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest, session string) (*mozi.AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse, error) {
	var resp mozi.AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
