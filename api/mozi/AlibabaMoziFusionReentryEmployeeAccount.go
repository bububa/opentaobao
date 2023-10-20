package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozifusionreentryemployeeaccount 重新入职并且重新启用账号
// alibaba.mozi.fusion.reentry.employee.account
//
// 重新入职并且重新启用账号
func Alibabamozifusionreentryemployeeaccount(clt *core.SDKClient, req *mozi.AlibabamozifusionreentryemployeeaccountAPIRequest, session string) (*mozi.AlibabamozifusionreentryemployeeaccountAPIResponse, error) {
	var resp mozi.AlibabamozifusionreentryemployeeaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
