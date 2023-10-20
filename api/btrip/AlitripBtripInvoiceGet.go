package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripinvoiceget 获取用户可用发票列表
// alitrip.btrip.invoice.get
//
// 差旅申请用户获取可用发票列表
func Alitripbtripinvoiceget(clt *core.SDKClient, req *btrip.AlitripbtripinvoicegetAPIRequest, session string) (*btrip.AlitripbtripinvoicegetAPIResponse, error) {
	var resp btrip.AlitripbtripinvoicegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
