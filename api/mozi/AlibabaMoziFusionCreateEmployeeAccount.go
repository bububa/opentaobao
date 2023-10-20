package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozifusioncreateemployeeaccount 创建MOZI自建人员和账号
// alibaba.mozi.fusion.create.employee.account
//
// 创建MOZI自建人员和账号
func Alibabamozifusioncreateemployeeaccount(clt *core.SDKClient, req *mozi.AlibabamozifusioncreateemployeeaccountAPIRequest, session string) (*mozi.AlibabamozifusioncreateemployeeaccountAPIResponse, error) {
	var resp mozi.AlibabamozifusioncreateemployeeaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
