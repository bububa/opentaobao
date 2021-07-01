package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

/* AlibabaMoziVdsTenantApiServiceDismiss
MOZI解除组织主管服务
alibaba.mozi.vds.tenant.api.service.dismiss

解除组织主管 */
func AlibabaMoziVdsTenantApiServiceDismiss(clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServiceDismissAPIRequest, session string) (*mozi.AlibabaMoziVdsTenantApiServiceDismissAPIResponse, error) {
	var resp mozi.AlibabaMoziVdsTenantApiServiceDismissAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
