package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziBucAccountPageall 查询租户内内所有账号
// alibaba.mozi.buc.account.pageall
//
// 查询租户内内所有账号
func AlibabaMoziBucAccountPageall(clt *core.SDKClient, req *mozi.AlibabaMoziBucAccountPageallAPIRequest, resp *mozi.AlibabaMoziBucAccountPageallAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
