package mozi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// Alibabamozibucaccountlistaccountids 根据一批账号ID查询账号列表
// alibaba.mozi.buc.account.list.accountids
//
// 根据一批账号ID查询账号列表
func Alibabamozibucaccountlistaccountids(clt *core.SDKClient, req *mozi.AlibabamozibucaccountlistaccountidsAPIRequest, session string) (*mozi.AlibabamozibucaccountlistaccountidsAPIResponse, error) {
	var resp mozi.AlibabamozibucaccountlistaccountidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
