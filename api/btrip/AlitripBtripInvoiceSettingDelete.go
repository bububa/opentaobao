package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripinvoicesettingdelete 发票删除
// alitrip.btrip.invoice.setting.delete
//
// 发票删除
func Alitripbtripinvoicesettingdelete(clt *core.SDKClient, req *btrip.AlitripbtripinvoicesettingdeleteAPIRequest, session string) (*btrip.AlitripbtripinvoicesettingdeleteAPIResponse, error) {
	var resp btrip.AlitripbtripinvoicesettingdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
