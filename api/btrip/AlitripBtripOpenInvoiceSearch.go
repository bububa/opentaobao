package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopeninvoicesearch 差旅申请用户搜索可用发票列表
// alitrip.btrip.open.invoice.search
//
// 差旅申请用户搜索可用发票列表
func Alitripbtripopeninvoicesearch(clt *core.SDKClient, req *btrip.AlitripbtripopeninvoicesearchAPIRequest, session string) (*btrip.AlitripbtripopeninvoicesearchAPIResponse, error) {
	var resp btrip.AlitripbtripopeninvoicesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
