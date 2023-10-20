package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozivdstenantapiservicematchempcodes 校验组-员工是否匹配
// alibaba.mozi.vds.tenant.api.service.matchempcodes
//
// 校验组-员工是否匹配
func Alibabamozivdstenantapiservicematchempcodes(clt *core.SDKClient, req *mozi.AlibabamozivdstenantapiservicematchempcodesAPIRequest, session string) (*mozi.AlibabamozivdstenantapiservicematchempcodesAPIResponse, error) {
	var resp mozi.AlibabamozivdstenantapiservicematchempcodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
