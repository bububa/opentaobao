package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozifusionupdateemployeeaccount 更新人员和账号属性
// alibaba.mozi.fusion.update.employee.account
//
// 更新人员和账号基本属性
func Alibabamozifusionupdateemployeeaccount(clt *core.SDKClient, req *mozi.AlibabamozifusionupdateemployeeaccountAPIRequest, session string) (*mozi.AlibabamozifusionupdateemployeeaccountAPIResponse, error) {
	var resp mozi.AlibabamozifusionupdateemployeeaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
