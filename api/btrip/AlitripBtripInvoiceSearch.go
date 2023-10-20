package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripinvoicesearch 根据发票抬头搜索发票
// alitrip.btrip.invoice.search
//
// 用户根据发票抬头搜索发票信息
func Alitripbtripinvoicesearch(clt *core.SDKClient, req *btrip.AlitripbtripinvoicesearchAPIRequest, session string) (*btrip.AlitripbtripinvoicesearchAPIResponse, error) {
	var resp btrip.AlitripbtripinvoicesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
