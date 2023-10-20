package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziBucAccountPageall 查询租户内内所有账号
// alibaba.mozi.buc.account.pageall
//
// 查询租户内内所有账号
func AlibabaMoziBucAccountPageall(clt *core.SDKClient, req *mozi.AlibabaMoziBucAccountPageallAPIRequest, session string) (*mozi.AlibabaMoziBucAccountPageallAPIResponse, error) {
	var resp mozi.AlibabaMoziBucAccountPageallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
