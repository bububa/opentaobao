package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripinvoicesettingrule 发票规则设置
// alitrip.btrip.invoice.setting.rule
//
// 发票规则设置
func Alitripbtripinvoicesettingrule(clt *core.SDKClient, req *btrip.AlitripbtripinvoicesettingruleAPIRequest, session string) (*btrip.AlitripbtripinvoicesettingruleAPIResponse, error) {
	var resp btrip.AlitripbtripinvoicesettingruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
