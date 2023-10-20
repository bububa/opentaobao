package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozifusionaddorupdateemployeeaccount 添加人员和账号复合接口
// alibaba.mozi.fusion.addorupdate.employee.account
//
// 添加人员和账号复合接口
func Alibabamozifusionaddorupdateemployeeaccount(clt *core.SDKClient, req *mozi.AlibabamozifusionaddorupdateemployeeaccountAPIRequest, session string) (*mozi.AlibabamozifusionaddorupdateemployeeaccountAPIResponse, error) {
	var resp mozi.AlibabamozifusionaddorupdateemployeeaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
