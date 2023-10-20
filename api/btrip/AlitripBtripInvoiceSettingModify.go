package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripinvoicesettingmodify 发票变更
// alitrip.btrip.invoice.setting.modify
//
// 发票变更
func Alitripbtripinvoicesettingmodify(clt *core.SDKClient, req *btrip.AlitripbtripinvoicesettingmodifyAPIRequest, session string) (*btrip.AlitripbtripinvoicesettingmodifyAPIResponse, error) {
	var resp btrip.AlitripbtripinvoicesettingmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
