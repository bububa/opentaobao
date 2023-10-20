package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozivdstenantapiservicetenantbyid 按租户ID查询租户信息
// alibaba.mozi.vds.tenant.api.service.tenantbyid
//
// 按租户ID查询租户信息
func Alibabamozivdstenantapiservicetenantbyid(clt *core.SDKClient, req *mozi.AlibabamozivdstenantapiservicetenantbyidAPIRequest, session string) (*mozi.AlibabamozivdstenantapiservicetenantbyidAPIResponse, error) {
	var resp mozi.AlibabamozivdstenantapiservicetenantbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
