package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozifusiondimissionemployeeaccount 人员离职
// alibaba.mozi.fusion.dimission.employee.account
//
// 人员离职并且回收账号
func Alibabamozifusiondimissionemployeeaccount(clt *core.SDKClient, req *mozi.AlibabamozifusiondimissionemployeeaccountAPIRequest, session string) (*mozi.AlibabamozifusiondimissionemployeeaccountAPIResponse, error) {
	var resp mozi.AlibabamozifusiondimissionemployeeaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
