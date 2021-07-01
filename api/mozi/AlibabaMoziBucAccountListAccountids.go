package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

/* AlibabaMoziBucAccountListAccountids
根据一批账号ID查询账号列表
alibaba.mozi.buc.account.list.accountids

根据一批账号ID查询账号列表 */
func AlibabaMoziBucAccountListAccountids(clt *core.SDKClient, req *mozi.AlibabaMoziBucAccountListAccountidsAPIRequest, session string) (*mozi.AlibabaMoziBucAccountListAccountidsAPIResponse, error) {
	var resp mozi.AlibabaMoziBucAccountListAccountidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
