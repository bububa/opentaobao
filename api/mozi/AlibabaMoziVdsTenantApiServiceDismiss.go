package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozivdstenantapiservicedismiss MOZI解除组织主管服务
// alibaba.mozi.vds.tenant.api.service.dismiss
//
// 解除组织主管
func Alibabamozivdstenantapiservicedismiss(clt *core.SDKClient, req *mozi.AlibabamozivdstenantapiservicedismissAPIRequest, session string) (*mozi.AlibabamozivdstenantapiservicedismissAPIResponse, error) {
	var resp mozi.AlibabamozivdstenantapiservicedismissAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
