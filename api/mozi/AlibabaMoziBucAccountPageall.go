package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozibucaccountpageall 查询租户内内所有账号
// alibaba.mozi.buc.account.pageall
//
// 查询租户内内所有账号
func Alibabamozibucaccountpageall(clt *core.SDKClient, req *mozi.AlibabamozibucaccountpageallAPIRequest, session string) (*mozi.AlibabamozibucaccountpageallAPIResponse, error) {
	var resp mozi.AlibabamozibucaccountpageallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
