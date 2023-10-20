package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadaccountbalanceget 查询账户余额
// alibaba.scbp.ad.account.balance.get
//
// 查询推广账户余额
func Alibabascbpadaccountbalanceget(clt *core.SDKClient, req *scbp.AlibabascbpadaccountbalancegetAPIRequest, session string) (*scbp.AlibabascbpadaccountbalancegetAPIResponse, error) {
	var resp scbp.AlibabascbpadaccountbalancegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
