package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziBucAccountListAccountids 根据一批账号ID查询账号列表
// alibaba.mozi.buc.account.list.accountids
//
// 根据一批账号ID查询账号列表
func AlibabaMoziBucAccountListAccountids(clt *core.SDKClient, req *mozi.AlibabaMoziBucAccountListAccountidsAPIRequest, resp *mozi.AlibabaMoziBucAccountListAccountidsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
