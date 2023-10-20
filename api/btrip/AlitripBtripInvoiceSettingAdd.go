package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripinvoicesettingadd 发票设置
// alitrip.btrip.invoice.setting.add
//
// 发票设置
func Alitripbtripinvoicesettingadd(clt *core.SDKClient, req *btrip.AlitripbtripinvoicesettingaddAPIRequest, session string) (*btrip.AlitripbtripinvoicesettingaddAPIResponse, error) {
	var resp btrip.AlitripbtripinvoicesettingaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
